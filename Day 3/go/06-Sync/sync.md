# Go 语言 `sync` 包

`sync` 包提供了基础的同步原语，用于解决并发编程中的竞态问题。以下是核心组件的介绍。

---

## 1. `sync.Mutex` — 互斥锁

最基础的锁，同一时刻只允许一个 goroutine 访问共享资源。

```go
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    count++
}
```

## 2. `sync.RWMutex` — 读写锁

允许多个读者同时访问，但写操作是排他的。适合**读多写少**的场景。

```go
var rw sync.RWMutex
var cache = make(map[string]string)

func read(key string) string {
    rw.RLock()
    defer rw.RUnlock()
    return cache[key]
}

func write(key, val string) {
    rw.Lock()
    defer rw.Unlock()
    cache[key] = val
}
```

## 3. `sync.WaitGroup` — 等待组

等待一组 goroutine 全部完成，常用于并发任务编排。

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("task %d done\n", id)
    }(i)
}
wg.Wait()
```

## 4. `sync.Once` — 仅执行一次

保证某个操作只执行一次，典型场景是**单例初始化**。

```go
var once sync.Once
var instance *DB

func GetDB() *DB {
    once.Do(func() {
        instance = &DB{...}
    })
    return instance
}
```

## 5. `sync.Map` — 并发安全的 Map

针对特定场景优化的并发 map，适合**读多写少**或**key 稳定**的场景。

```go
var m sync.Map

m.Store("key", "value")
val, ok := m.Load("key")
m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true // 返回 false 停止遍历
})
```

> ⚠️ 一般场景下，`map` + `Mutex` 的性能往往更好。`sync.Map` 不是万能的。

## 6. `sync.Pool` — 对象池

缓存和复用临时对象，减少 GC 压力。典型场景是高频分配/回收的场景（如 HTTP 请求处理）。

```go
var bufPool = sync.Pool{
    New: func() any {
        return new(bytes.Buffer)
    },
}

func handler() {
    buf := bufPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufPool.Put(buf)
    // 使用 buf...
}
```

## 7. `sync.Cond` — 条件变量

用于 goroutine 之间的条件等待和通知，适合**生产者-消费者**模型。

```go
var mu sync.Mutex
var cond = sync.NewCond(&mu)
var queue []int

// 消费者
go func() {
    cond.L.Lock()
    for len(queue) == 0 {
        cond.Wait() // 释放锁并等待
    }
    item := queue[0]
    queue = queue[1:]
    cond.L.Unlock()
}()

// 生产者
cond.L.Lock()
queue = append(queue, 42)
cond.Signal() // 唤醒一个等待的 goroutine
cond.L.Unlock()
```

---

## 速查表

| 类型 | 用途 | 关键方法 |
|---|---|---|
| `Mutex` | 互斥锁 | `Lock()`, `Unlock()` |
| `RWMutex` | 读写锁 | `RLock()`, `RUnlock()`, `Lock()`, `Unlock()` |
| `WaitGroup` | 等待并发完成 | `Add()`, `Done()`, `Wait()` |
| `Once` | 单次执行 | `Do(func)` |
| `Map` | 并发安全 map | `Store()`, `Load()`, `Delete()`, `Range()` |
| `Pool` | 对象池 | `Get()`, `Put()` |
| `Cond` | 条件等待/通知 | `Wait()`, `Signal()`, `Broadcast()` |

---

## 使用建议

1. **优先用 channel**：Go 的哲学是"通过通信共享内存"，`sync` 包适合底层或性能敏感场景
2. **锁的粒度要小**：`Lock()` 和 `Unlock()` 之间的代码尽量少
3. **避免嵌套锁**：容易造成死锁
4. **`defer Unlock()`**：防止提前 return 导致锁未释放

有具体的使用场景想深入讨论的话，随时问我。