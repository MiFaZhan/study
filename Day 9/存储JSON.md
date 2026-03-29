**Redis 没有原生的 JSON 数据类型**（传统 Redis）。

在这个代码中，JSON 是作为 **String 类型** 存储的。

## Redis 的 5 种基础数据类型

```
┌─────────────────────────────────────────────────────┐
│  Redis 数据类型                                       │
│                                                     │
│  1. String  ← JSON 字符串存在这里                    │
│  2. Hash     (字段-值对)                              │
│  3. List     (有序列表)                              │
│  4. Set      (无序集合)                              │
│  5. Sorted Set (有序集合，带分数)                     │
└─────────────────────────────────────────────────────┘
```

## 实际存储过程

```go
// 1. Go 结构体
product := &Product{ID: 123, Name: "iPhone", Price: 6999}

// 2. 序列化成 JSON 字符串
data := json.Marshal(product)
// 结果: {"id":123,"name":"iPhone","price":6999}

// 3. 作为 String 存入 Redis
rdb.Set(ctx, "product:123", data, 10*time.Minute)
// Redis 中: key="product:123", value="{"id":123,"name":"iPhone","price":6999}"
```

**Redis CLI 中查看**：
```bash
127.0.0.1:6379> GET product:123
"{"id":123,"name":"iPhone","price":6999}"  ← 这是一个 String 类型
```

## Redis 7.0 的 JSON 模块（可选）

Redis 7.0 之后引入了 JSON 模块，支持原生的 JSON 操作：

```bash
# 安装 JSON 模块后
JSON.SET product:123 $ '{"id":123,"name":"iPhone"}'
JSON.GET product:123
JSON.GET product:123 $.name  # 只获取 name 字段
```

**但大多数项目还是用 String 存储 JSON**，因为：
- 兼容性好（所有 Redis 版本都支持）
- 简单直接
- 不需要额外安装模块

## 其他存储方式

除了 String 存 JSON，还可以用 **Hash**：

```go
// 用 Hash 存储商品信息
rdb.HSet(ctx, "product:123",
    "id", 123,
    "name", "iPhone",
    "price", 6999,
    "stock", 100,
)

// 读取
rdb.HGetAll(ctx, "product:123")
```

**对比**：
| 方式 | 优点 | 缺点 |
|------|------|------|
| **String + JSON** | 简单，跨语言 | 无法部分更新 |
| **Hash** | 可以单独修改某个字段 | 结构固定，不如 JSON 灵活 |

所以这个代码用的是 **String 类型存储 JSON 字符串**，这是最常见的做法。
        