# Go 语言内存管理与 GC 机制

## 一、内存管理架构

Go 的内存管理采用**多级缓存**的设计，核心分为四层：

```
┌─────────────────────────────────────┐
│           用户代码 (malloc / make)      │
├─────────────────────────────────────┤
│          mcache (每 P 独有)             │  ← 无锁，快速分配
├─────────────────────────────────────┤
│          mcentral (全局, 每 sizeclass)   │  ← P 间共享
├─────────────────────────────────────┤
│          mheap (全局堆)                 │  ← 向 OS 申请内存
├─────────────────────────────────────┤
│          操作系统 (mmap / brk)          │
└─────────────────────────────────────┘
```

### 1. mcache — 线程本地缓存

每个逻辑处理器 **P** 拥有一个独立的 mcache，分配小对象时**无需加锁**：

```go
// 每个 mcache 维护 ~70 个 sizeclass 的 mspan
type mcache struct {
    alloc [numSpanClasses]*mspan
}
```

### 2. mcentral — 中心缓存

按 sizeclass 组织，当 mcache 用完时从 mcentral 补充（需要加锁）：

```go
type mcentral struct {
    spanclass spanClass
    partial  [2]spanSet   // 有空闲对象的 span
    full     [2]spanSet   // 无空闲对象的 span
}
```

### 3. mheap — 堆管理

- 以 **8KB** 的页（page）为基本单位
- 通过 **radix tree** 管理所有页的元数据
- 向操作系统申请内存（Linux 上用 `mmap`）

### 4. mspan — 内存块

Go 把内存按 **sizeclass**（共 ~70 种）分组，每种 span 固定分配特定大小的对象：

| sizeclass | 对象大小 | span 大小 | 对象数量 |
|-----------|---------|----------|---------|
| 1         | 8B      | 8KB      | 1024    |
| 5         | 48B     | 8KB      | 170     |
| 10        | 128B    | 8KB      | 64      |
| ...       | ...     | ...      | ...     |

### 分配流程

```
make([]byte, 64)
    │
    ├── ≤ 32KB (小对象)
    │   ├── 确定 sizeclass
    │   ├── mcache.alloc[sizeclass] 有空闲 → 直接返回（无锁）
    │   └── 无空闲 → 从 mcentral 补充 → 再分配
    │
    ├── 32KB ~ 32MB (大对象)
    │   └── 直接从 mheap 分配 span
    │
    └── > 32MB (巨大对象)
        └── 直接向 OS 申请
```

**微小对象（< 16B）** 还会走 **tiny allocator**，将多个 tiny 对象合并到一个 slot 中，进一步减少分配开销。

---

## 二、GC 机制 — 三色并发标记清除

Go 使用的是 **三色标记-清除（Tri-color Mark-and-Sweep）** 算法，从 Go 1.8 起支持**并发清除**，STW（Stop-The-World）时间通常在 **亚毫秒级**。

### 核心算法

```
初始状态：所有对象白色

┌──────────────────────────────────────────┐
│  1. STW：将 GC Root 置灰（很短暂）          │
│  2. STW 结束，开始并发标记                   │
│  3. 从灰色对象出发：                         │
│     - 灰色 → 扫描其引用的白色对象 → 置灰      │
│     - 灰色 → 扫描完毕 → 置黑                │
│  4. 无灰色对象时：剩余白色 = 垃圾             │
│  5. 并发清除白色对象                         │
└──────────────────────────────────────────┘
```

```
    ┌──────┐     ┌──────┐     ┌──────┐
    │ 白色  │────▶│ 灰色  │────▶│ 黑色  │
    │ 垃圾  │     │ 待扫描 │     │ 存活  │
    └──────┘     └──────┘     └──────┘
      ↑                          │
      └──────── 清除 ◀────────────┘
```

### 三色不变式

要保证并发标记的正确性，必须维持以下任一不变式：

**强三色不变式**：黑色对象不能直接引用白色对象

**弱三色不变式**：黑色对象引用白色对象时，该白色对象必须有灰色祖先

### 写屏障（Write Barrier）

Go 1.8+ 使用 **混合写屏障（Hybrid Write Barrier）**，结合了插入写屏障和删除写屏障的优点：

```go
// 混合写屏障的伪代码
func writePointer(slot *unsafe.Pointer, ptr unsafe.Pointer) {
    // 1. 插入写屏障部分：新引用的对象标灰
    shade(ptr)
    // 2. 删除写屏障部分：旧引用的对象标灰
    shade(*slot)
    // 3. 实际写入
    *slot = ptr
}
```

混合写屏障使得 Go 1.8+ 的 STW 时间从之前的 ~1ms 降到 ~100μs 以下。

---

## 三、GC 触发时机

```go
// 由 runtime 控制，三种触发方式：

// 1. 堆内存达到阈值（默认为上次 GC 后堆大小的 2 倍）
//    可通过 GOGC 调节（默认 100，即堆增长 100% 时触发）
//    Go 1.19+ 引入 GOMEMLIMIT（软内存上限）

// 2. 定时触发（默认 2 分钟，防止长时间不 GC）

// 3. 手动触发
runtime.GC()
```

**GOGC 与 GOMEMLIMIT**：

```bash
# 传统方式：GOGC=100 表示当新分配的堆内存 = 上次 GC 后存活量时触发
GOGC=100 ./myapp

# Go 1.19+：设置软内存上限（更适合容器环境）
GOMEMLIMIT=512MiB ./myapp
GOMEMLIMIT=512MiB GOGC=off ./myapp  # 关闭 GOGC，仅靠内存上限
```

---

## 四、GC 的五个阶段

```
┌─────────────────────────────────────────────────────┐
│  阶段          | STW? | 说明                         │
├─────────────────────────────────────────────────────┤
│  1. GCMark     | 是    | 开启写屏障，标记根对象（~μs）  │
│  2. GCMark     | 否    | 并发标记，与用户代码并行        │
│  3. GCMarkTerm | 是    | 重新扫描部分根对象（~μs）      │
│  4. GCSweep    | 否    | 并发清除，归还内存             │
│  5. GCSweepTerm| 是    | 清除结束（极短）              │
└─────────────────────────────────────────────────────┘
```

---

## 五、GC 辅助机制

### 1. Mark Assist（标记辅助）

当用户 goroutine 分配速度过快，超过 GC 标记速度时，**分配内存的 goroutine 被要求帮忙标记**：

```go
// 伪代码
func mallocgc(size uintptr) unsafe.Pointer {
    if gcphase == _GCmark && assistRequired {
        gcAssistAlloc()  // 这个 goroutine 帮忙做标记工作
    }
    // ... 正常分配
}
```

### 2. Sweeper（清除器）

清除阶段与用户代码并发执行，逐个归还 span 中的空闲对象到 mcentral/mcache。

### 3. 内存归还策略

```go
// Go 不会立即将内存归还 OS，而是：
// 1. 每 5 分钟检查一次空闲 span
// 2. 连续 5 分钟空闲的内存通过 madvise(MADV_DONTNEED) 归还 OS
//    （物理内存释放，虚拟地址保留）

// Go 1.16+ 引入更积极的内存归还策略
// 可通过 debug.FreeOSMemory() 手动触发
```

---

## 六、逃逸分析

Go 编译器在编译期进行**逃逸分析**，决定对象分配在栈上还是堆上：

```go
// 栈分配 — 函数返回后自动回收
func stackAlloc() int {
    x := 42      // 栈上分配
    return x
}

// 堆分配 — 指针逃逸到函数外部
func heapAlloc() *int {
    x := 42      // 逃逸到堆上
    return &x
}

// 用 -gcflags="-m" 查看逃逸分析结果
// go build -gcflags="-m" main.go
```

**常见逃逸场景**：
- 返回局部变量的指针
- 将局部变量发送到 channel
- 赋值给 `interface{}` 类型
- 切片/字典扩容导致底层数组逃逸
- 闭包引用外部变量

---

## 七、性能调优实践

```go
// 1. 减少堆分配
//    - 优先使用值类型而非指针
//    - 复用对象（sync.Pool）

var bufPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 4096)
    },
}

func handle() {
    buf := bufPool.Get().([]byte)
    defer bufPool.Put(buf[:0])
    // 使用 buf...
}

// 2. 预分配容量
m := make(map[string]int, 1000)  // 避免多次扩容
s := make([]int, 0, 1000)         // 避免多次扩容

// 3. 用 GOGC/GOMEMLIMIT 调节 GC 频率
//    内存充足但希望减少 GC → 调大 GOGC
//    容器有内存限制 → 设置 GOMEMLIMIT
```

---

## 八、总结

| 特性 | Go 的实现 |
|------|----------|
| 分配器 | TCMalloc 风格，多级缓存 |
| GC 算法 | 三色并发标记清除 |
| 写屏障 | 混合写屏障（Go 1.8+） |
| STW 时间 | 亚毫秒级（通常 < 100μs） |
| 触发方式 | GOGC 阈值 / 定时 / 手动 |
| 内存归还 | 延迟释放，通过 madvise |
| 调优手段 | sync.Pool、逃逸分析、GOGC/GOMEMLIMIT |

Go 的 GC 设计哲学是在**延迟、吞吐、内存开销**三者之间取平衡，牺牲一定的 CPU 和内存利用率，换取极低的暂停时间和简单易用的编程体验。