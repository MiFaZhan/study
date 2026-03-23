好的，这是一个非常实用的技术话题。我将为你详细讲解 Redis 的各种数据类型及其操作，并配以 Go 语言（Golang）的对应实现代码。

---

### 第一部分：Redis 核心数据类型与操作

Redis 不仅仅是一个简单的键值存储，它支持丰富的数据结构。以下是其核心类型及常用操作命令。

#### 1. String (字符串)
最基础的类型，可以包含任何数据（文本、数字、序列化对象），最大 512MB。
*   **典型用例**： 缓存、计数器、分布式锁。
*   **常用命令**：
    ```bash
    # 设置值（SET）与获取值（GET）
    SET mykey "Hello Redis"
    GET mykey  # 返回 "Hello Redis"

    # 设置值并指定过期时间（秒）
    SET mykey "Hello" EX 10

    # 仅当键不存在时设置（常用于实现锁）
    SET mykey "Hello" NX
    # 仅当键存在时设置
    SET mykey "World" XX

    # 自增/自减（用于计数器）
    SET counter 100
    INCR counter  # 返回 101
    INCRBY counter 10  # 返回 111
    DECR counter  # 返回 110

    # 批量设置/获取
    MSET key1 "val1" key2 "val2"
    MGET key1 key2  # 返回列表 ["val1", "val2"]
    ```

#### 2. List (列表)
有序的字符串列表，按照插入顺序排序。底层是双向链表，头部和尾部操作极快。
*   **典型用例**： 消息队列、最新消息排行（如微博时间线）、栈。
*   **常用命令**：
    ```bash
    # 从左端（头部）推入
    LPUSH mylist "world"
    LPUSH mylist "hello"  # 列表现在是 ["hello", "world"]

    # 从右端（尾部）推入
    RPUSH mylist "!"  # 列表现在是 ["hello", "world", "!"]

    # 从左端弹出
    LPOP mylist  # 返回 "hello"

    # 从右端弹出
    RPOP mylist  # 返回 "!"

    # 获取列表长度
    LLEN mylist

    # 获取指定范围的元素 (0 到 -1 表示所有元素)
    LRANGE mylist 0 -1

    # 阻塞式弹出（用于可靠队列）
    BLPOP mylist 30  # 阻塞30秒等待元素
    ```

#### 3. Hash (哈希)
一个键值对集合，适合存储对象。
*   **典型用例**： 存储用户信息、商品信息等结构化数据。
*   **常用命令**：
    ```bash
    # 设置单个字段
    HSET user:1000 name "John Doe"
    HSET user:1000 email "john@example.com"

    # 获取单个字段
    HGET user:1000 name  # 返回 "John Doe"

    # 设置多个字段
    HMSET user:1000 name "Jane Doe" age 30 city "New York"

    # 获取所有字段和值
    HGETALL user:1000

    # 获取所有字段
    HKEYS user:1000

    # 获取所有值
    HVALS user:1000

    # 字段自增
    HINCRBY user:1000 age 1  # 返回 31

    # 检查字段是否存在
    HEXISTS user:1000 email  # 返回 1 (true)
    ```

#### 4. Set (集合)
无序的字符串集合，元素唯一。
*   **典型用例**： 标签、共同好友（交集）、唯一性统计。
*   **常用命令**：
    ```bash
    # 添加元素
    SADD myset "apple"
    SADD myset "banana" "cherry"

    # 获取所有元素
    SMEMBERS myset  # 返回顺序不定

    # 判断元素是否存在
    SISMEMBER myset "apple"  # 返回 1 (true)

    # 获取集合大小
    SCARD myset

    # 集合运算：交集、并集、差集
    SADD set1 "a" "b" "c"
    SADD set2 "b" "c" "d"
    SINTER set1 set2  # 返回 {"b", "c"}
    SUNION set1 set2  # 返回 {"a","b","c","d"}
    SDIFF set1 set2   # 返回 {"a"} (在set1但不在set2)

    # 随机获取元素
    SRANDMEMBER myset 2  # 随机返回2个不重复元素
    SPOP myset  # 随机移除并返回一个元素
    ```

#### 5. Sorted Set (有序集合 / ZSet)
与 Set 类似，但每个元素都会关联一个 `double` 类型的分数（score）。Redis 通过分数来为集合中的成员进行从小到大的排序。
*   **典型用例**： 排行榜、带权重的任务队列。
*   **常用命令**：
    ```bash
    # 添加元素（需指定分数）
    ZADD leaderboard 100 "player:1"
    ZADD leaderboard 250 "player:2" 150 "player:3"

    # 获取指定范围的元素（按分数从小到大）
    ZRANGE leaderboard 0 -1  # 返回 ["player:1", "player:3", "player:2"]
    # 带上分数
    ZRANGE leaderboard 0 -1 WITHSCORES

    # 获取指定范围的元素（按分数从大到小）
    ZREVRANGE leaderboard 0 1  # 返回前两名 ["player:2", "player:3"]

    # 获取元素的分数
    ZSCORE leaderboard "player:1"  # 返回 100

    # 获取元素的排名（从0开始，按分数从小到大）
    ZRANK leaderboard "player:2"  # 返回 2 (第一名)
    # 按分数从大到小排名
    ZREVRANK leaderboard "player:2"  # 返回 0 (第一名)

    # 对分数进行自增
    ZINCRBY leaderboard 50 "player:1"  # player:1 的分数变为150

    # 获取指定分数范围内的元素数量
    ZCOUNT leaderboard 100 200

    # 集合运算（交集、并集）并存储结果
    ZINTERSTORE out_set 2 set1 set2 WEIGHTS 2 3  # 复杂，需指定权重和聚合方式
    ```

#### 6. 其他特殊类型 (了解即可)
*   **Bitmap (位图)**： 不是独立数据类型，而是基于 String 类型的一组面向位的操作。非常适合布尔型统计（如用户签到）。
    ```bash
    SETBIT sign:20231001 1001 1  # 用户1001在2023-10-01签到
    GETBIT sign:20231001 1001    # 返回 1
    BITCOUNT sign:20231001       # 统计当天签到总数
    ```
*   **HyperLogLog**： 用于基数估算（不重复元素个数），占用固定 12KB 内存。
    ```bash
    PFADD visitors "ip1" "ip2" "ip1"
    PFCOUNT visitors  # 返回 2 (估算值)
    ```
*   **Stream (流)**： Redis 5.0 引入，类似 Kafka 的日志数据结构，用于消息队列和事件流。
    ```bash
    XADD mystream * sensor-id 1234 temperature 19.8
    XREAD COUNT 2 STREAMS mystream 0
    ```

---

### 第二部分：Golang 操作 Redis

在 Go 中，最流行和强大的 Redis 客户端库是 `github.com/redis/go-redis/v9`。下面以这个库为例，展示如何操作上述各种类型。

#### 1. 连接 Redis
```go
package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 地址
        Password: "",               // 密码，没有则留空
        DB:       0,                // 使用默认 DB
    })

    // 测试连接
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
}
```

#### 2. 操作 String 类型
```go
func stringOperations() {
    // SET & GET
    err := rdb.Set(ctx, "mykey", "Hello from Go", 10*time.Second).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "mykey").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("mykey:", val) // 输出: mykey: Hello from Go

    // SETNX (仅当 key 不存在)
    success, err := rdb.SetNX(ctx, "lock_key", "locked", 30*time.Second).Result()
    fmt.Println("SETNX succeeded:", success)

    // INCR (自增)
    rdb.Del(ctx, "counter") // 先删除旧值
    num, err := rdb.Incr(ctx, "counter").Result()
    fmt.Println("Counter:", num) // 输出: Counter: 1

    // MSET & MGET
    rdb.MSet(ctx, "k1", "v1", "k2", "v2")
    vals, _ := rdb.MGet(ctx, "k1", "k2").Result()
    fmt.Println("MGET:", vals) // 输出: MGET: [v1 v2]
}
```

#### 3. 操作 List 类型
```go
func listOperations() {
    key := "mylist"
    rdb.Del(ctx, key) // 清空

    // LPUSH & RPUSH
    rdb.LPush(ctx, key, "world")
    rdb.LPush(ctx, key, "hello")
    rdb.RPush(ctx, key, "!")

    // LRANGE
    list, _ := rdb.LRange(ctx, key, 0, -1).Result()
    fmt.Println("List:", list) // 输出: List: [hello world !]

    // LPOP & RPOP
    left, _ := rdb.LPop(ctx, key).Result()
    fmt.Println("LPOP:", left) // 输出: LPOP: hello
    right, _ := rdb.RPop(ctx, key).Result()
    fmt.Println("RPOP:", right) // 输出: RPOP: !

    // LLEN
    length, _ := rdb.LLen(ctx, key).Result()
    fmt.Println("Length:", length) // 输出: Length: 1

    // BLPOP (阻塞弹出)
    go func() {
        time.Sleep(2 * time.Second)
        rdb.RPush(ctx, "blkey", "task1")
    }()
    result, _ := rdb.BLPop(ctx, 5*time.Second, "blkey").Result()
    fmt.Println("BLPOP result:", result) // 输出: BLPOP result: [blkey task1]
}
```

#### 4. 操作 Hash 类型
```go
func hashOperations() {
    key := "user:1001"
    rdb.Del(ctx, key)

    // HSET & HGET
    rdb.HSet(ctx, key, "name", "Alice")
    rdb.HSet(ctx, key, "email", "alice@example.com")
    name, _ := rdb.HGet(ctx, key, "name").Result()
    fmt.Println("Name:", name) // 输出: Name: Alice

    // HMSET (一次性设置多个字段)
    user := map[string]interface{}{
        "name":  "Bob",
        "email": "bob@example.com",
        "age":   25,
    }
    rdb.HMSet(ctx, key, user)

    // HGETALL
    all, _ := rdb.HGetAll(ctx, key).Result()
    fmt.Println("User info:", all) // 输出: User info: map[age:25 email:bob@example.com name:Bob]

    // HINCRBY
    age, _ := rdb.HIncrBy(ctx, key, "age", 1).Result()
    fmt.Println("New age:", age) // 输出: New age: 26

    // HEXISTS
    exists, _ := rdb.HExists(ctx, key, "email").Result()
    fmt.Println("Email exists:", exists) // 输出: Email exists: true
}
```

#### 5. 操作 Set 类型
```go
func setOperations() {
    key1 := "set1"
    key2 := "set2"
    rdb.Del(ctx, key1, key2)

    // SADD
    rdb.SAdd(ctx, key1, "a", "b", "c")
    rdb.SAdd(ctx, key2, "b", "c", "d")

    // SMEMBERS
    members, _ := rdb.SMembers(ctx, key1).Result()
    fmt.Println("Set1 members:", members) // 顺序不定

    // SISMEMBER
    isMember, _ := rdb.SIsMember(ctx, key1, "a").Result()
    fmt.Println("'a' is in set1:", isMember) // 输出: 'a' is in set1: true

    // SCARD
    count, _ := rdb.SCard(ctx, key1).Result()
    fmt.Println("Set1 cardinality:", count) // 输出: Set1 cardinality: 3

    // SINTER (交集)
    inter, _ := rdb.SInter(ctx, key1, key2).Result()
    fmt.Println("Intersection:", inter) // 输出: Intersection: [b c] (顺序不定)

    // SUNION (并集)
    union, _ := rdb.SUnion(ctx, key1, key2).Result()
    fmt.Println("Union:", union) // 输出: Union: [a b c d] (顺序不定)

    // SDIFF (差集)
    diff, _ := rdb.SDiff(ctx, key1, key2).Result()
    fmt.Println("Difference (set1 - set2):", diff) // 输出: Difference: [a]
}
```

#### 6. 操作 Sorted Set 类型
```go
func sortedSetOperations() {
    key := "leaderboard"
    rdb.Del(ctx, key)

    // ZADD
    members := []*redis.Z{
        {Score: 100, Member: "player:1"},
        {Score: 250, Member: "player:2"},
        {Score: 150, Member: "player:3"},
    }
    rdb.ZAdd(ctx, key, members...)

    // ZRANGE (从小到大)
    all, _ := rdb.ZRange(ctx, key, 0, -1).Result()
    fmt.Println("All players (asc):", all) // 输出: All players: [player:1 player:3 player:2]

    // ZRANGEWITHSCORES
    allWithScores, _ := rdb.ZRangeWithScores(ctx, key, 0, -1).Result()
    for _, z := range allWithScores {
        fmt.Printf("Player: %s, Score: %.0f\n", z.Member, z.Score)
    }

    // ZREVRANGE (从大到小)
    top2, _ := rdb.ZRevRange(ctx, key, 0, 1).Result()
    fmt.Println("Top 2 players:", top2) // 输出: Top 2 players: [player:2 player:3]

    // ZSCORE
    score, _ := rdb.ZScore(ctx, key, "player:1").Result()
    fmt.Println("Player:1 score:", score) // 输出: Player:1 score: 100

    // ZRANK & ZREVRANK
    rank, _ := rdb.ZRank(ctx, key, "player:2").Result() // 从小到大排名，从0开始
    revRank, _ := rdb.ZRevRank(ctx, key, "player:2").Result() // 从大到小排名
    fmt.Printf("Player:2 rank: %d (asc), %d (desc)\n", rank, revRank) // 输出: Player:2 rank: 2 (asc), 0 (desc)

    // ZINCRBY
    newScore, _ := rdb.ZIncrBy(ctx, key, 50, "player:1").Result()
    fmt.Println("Player:1 new score:", newScore) // 输出: Player:1 new score: 150

    // ZCOUNT
    count, _ := rdb.ZCount(ctx, "100", "200").Result()
    fmt.Println("Players with score between 100 and 200:", count) // 输出: ...: 2
}
```

#### 7. 通用操作与最佳实践
```go
func generalOperations() {
    // 检查键是否存在
    exists, _ := rdb.Exists(ctx, "mykey").Result()
    fmt.Println("Key exists:", exists) // 1 存在，0 不存在

    // 设置过期时间
    rdb.Expire(ctx, "mykey", 1*time.Hour)

    // 获取剩余生存时间 (TTL)
    ttl, _ := rdb.TTL(ctx, "mykey").Result()
    fmt.Println("TTL:", ttl) // 输出持续时间，-1 永不过期，-2 键不存在

    // 删除键
    rdb.Del(ctx, "key1", "key2")

    // 使用 Pipeline 批量执行命令，减少 RTT (往返时间)
    pipe := rdb.Pipeline()
    incr := pipe.Incr(ctx, "pipeline_counter")
    pipe.Expire(ctx, "pipeline_counter", time.Hour)
    _, err := pipe.Exec(ctx)
    if err != nil {
        panic(err)
    }
    fmt.Println("Pipeline result:", incr.Val())

    // 事务 (MULTI/EXEC)
    // 注意：Redis 事务不支持回滚，只是保证命令连续执行
    tx := rdb.TxPipeline()
    tx.Set(ctx, "txkey", "txval", 0)
    tx.Incr(ctx, "tx_counter")
    _, err = tx.Exec(ctx)
}
```

### 总结

| Redis 类型 | Go `go-redis` 方法 (常见) | 主要用途 |
| :--- | :--- | :--- |
| **String** | `Set`, `Get`, `Incr`, `MSet`, `MGet` | 缓存、计数器、分布式锁 |
| **List** | `LPush`, `RPush`, `LPop`, `RPop`, `LRange`, `BLPop` | 队列、栈、最新列表 |
| **Hash** | `HSet`, `HGet`, `HMSet`, `HGetAll`, `HIncrBy` | 存储对象/属性 |
| **Set** | `SAdd`, `SMembers`, `SIsMember`, `SInter`, `SUnion` | 标签、共同好友、去重 |
| **Sorted Set** | `ZAdd`, `ZRange`, `ZRevRange`, `ZScore`, `ZIncrBy`, `ZRank` | 排行榜、带权重队列 |
| **通用** | `Del`, `Exists`, `Expire`, `TTL`, `Pipeline`, `TxPipeline` | 键管理、批量操作、事务 |

**关键提示**：
1.  **错误处理**： 所有命令都返回 `error`，生产代码中必须检查。
2.  **上下文**： 使用 `context.Context` 来控制超时和取消。
3.  **连接池**： `go-redis` 自带连接池，`redis.Options` 中可以配置池大小等参数。
4.  **序列化**： 存储复杂结构体时，通常先用 `json.Marshal` 序列化为字符串再存入 Redis String 或 Hash。

希望这份详细的讲解和代码示例能帮助你熟练掌握 Redis 及其在 Go 中的应用！