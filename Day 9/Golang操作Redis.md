# Golang 操作 Redis 详解

## 一、安装 Redis 客户端库

Go 语言中最常用的 Redis 客户端是 `go-redis`。

```bash
go get github.com/redis/go-redis/v9
```

## 二、连接 Redis

### 1. 单机连接

```go
package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    // 创建 Redis 客户端
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",  // Redis 地址
        Password: "",                 // 密码，没有则为空
        DB:       0,                  // 使用默认数据库
    })

    // 测试连接
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("连接成功:", pong)
}
```

### 2. 连接池配置

```go
rdb := redis.NewClient(&redis.Options{
    Addr:         "localhost:6379",
    Password:     "",
    DB:           0,
    PoolSize:     10,              // 连接池大小
    MinIdleConns: 5,               // 最小空闲连接数
})
```


## 三、String 操作

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func StringOperations(rdb *redis.Client) {
    // 1. 设置值
    err := rdb.Set(ctx, "name", "张三", 0).Err()
    if err != nil {
        panic(err)
    }

    // 2. 获取值
    val, err := rdb.Get(ctx, "name").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("name:", val)

    // 3. 设置带过期时间的值
    err = rdb.Set(ctx, "session", "abc123", 10*time.Minute).Err()
    
    // 4. 仅当 key 不存在时设置
    success, err := rdb.SetNX(ctx, "lock", "locked", 30*time.Second).Result()
    fmt.Println("SetNX 成功:", success)

    // 5. 批量设置
    err = rdb.MSet(ctx, "key1", "value1", "key2", "value2").Err()

    // 6. 批量获取
    vals, err := rdb.MGet(ctx, "key1", "key2").Result()
    fmt.Println("MGet:", vals)

    // 7. 递增
    count, err := rdb.Incr(ctx, "counter").Result()
    fmt.Println("Counter:", count)

    // 8. 递增指定值
    count, err = rdb.IncrBy(ctx, "counter", 10).Result()
    fmt.Println("Counter after IncrBy:", count)

    // 9. 递减
    count, err = rdb.Decr(ctx, "counter").Result()

    // 10. 追加字符串
    length, err := rdb.Append(ctx, "name", "先生").Result()
    fmt.Println("字符串长度:", length)

    // 11. 获取字符串长度
    length, err = rdb.StrLen(ctx, "name").Result()

    // 12. 删除键
    err = rdb.Del(ctx, "key1", "key2").Err()
}
```

## 四、Hash 操作

```go
func HashOperations(rdb *redis.Client) {
    // 1. 设置单个字段
    err := rdb.HSet(ctx, "user:1", "name", "张三").Err()
    
    // 2. 设置多个字段
    err = rdb.HSet(ctx, "user:1", map[string]interface{}{
        "name":  "张三",
        "age":   25,
        "email": "zhangsan@example.com",
    }).Err()

    // 3. 获取单个字段
    name, err := rdb.HGet(ctx, "user:1", "name").Result()
    fmt.Println("Name:", name)

    // 4. 获取多个字段
    vals, err := rdb.HMGet(ctx, "user:1", "name", "age").Result()
    fmt.Println("HMGet:", vals)

    // 5. 获取所有字段和值
    all, err := rdb.HGetAll(ctx, "user:1").Result()
    fmt.Println("HGetAll:", all)

    // 6. 获取所有字段名
    keys, err := rdb.HKeys(ctx, "user:1").Result()
    fmt.Println("HKeys:", keys)

    // 7. 获取所有值
    values, err := rdb.HVals(ctx, "user:1").Result()
    fmt.Println("HVals:", values)

    // 8. 判断字段是否存在
    exists, err := rdb.HExists(ctx, "user:1", "name").Result()
    fmt.Println("HExists:", exists)

    // 9. 删除字段
    err = rdb.HDel(ctx, "user:1", "email").Err()

    // 10. 获取字段数量
    count, err := rdb.HLen(ctx, "user:1").Result()
    fmt.Println("HLen:", count)

    // 11. 字段值递增
    newAge, err := rdb.HIncrBy(ctx, "user:1", "age", 1).Result()
    fmt.Println("New Age:", newAge)

    // 12. 仅当字段不存在时设置
    success, err := rdb.HSetNX(ctx, "user:1", "phone", "13800138000").Result()
    fmt.Println("HSetNX 成功:", success)
}
```

## 五、List 操作

```go
func ListOperations(rdb *redis.Client) {
    // 1. 从左侧插入
    err := rdb.LPush(ctx, "tasks", "task1", "task2").Err()

    // 2. 从右侧插入
    err = rdb.RPush(ctx, "tasks", "task3", "task4").Err()

    // 3. 从左侧弹出
    val, err := rdb.LPop(ctx, "tasks").Result()
    fmt.Println("LPop:", val)

    // 4. 从右侧弹出
    val, err = rdb.RPop(ctx, "tasks").Result()
    fmt.Println("RPop:", val)

    // 5. 获取列表长度
    length, err := rdb.LLen(ctx, "tasks").Result()
    fmt.Println("List Length:", length)

    // 6. 获取指定范围的元素
    vals, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
    fmt.Println("LRange:", vals)

    // 7. 获取指定索引的元素
    val, err = rdb.LIndex(ctx, "tasks", 0).Result()
    fmt.Println("LIndex:", val)

    // 8. 设置指定索引的值
    err = rdb.LSet(ctx, "tasks", 0, "new_task").Err()

    // 9. 在指定元素前插入
    err = rdb.LInsertBefore(ctx, "tasks", "task2", "task1.5").Err()

    // 10. 删除指定数量的元素
    err = rdb.LRem(ctx, "tasks", 1, "task1").Err()

    // 11. 保留指定范围的元素
    err = rdb.LTrim(ctx, "tasks", 0, 9).Err()

    // 12. 阻塞式弹出
    result, err := rdb.BLPop(ctx, 10*time.Second, "tasks").Result()
    if err != nil {
        fmt.Println("BLPop timeout")
    } else {
        fmt.Println("BLPop:", result)
    }
}
```


## 六、Set 操作

```go
func SetOperations(rdb *redis.Client) {
    // 1. 添加元素
    err := rdb.SAdd(ctx, "tags", "redis", "database", "cache").Err()

    // 2. 获取所有元素
    members, err := rdb.SMembers(ctx, "tags").Result()
    fmt.Println("SMembers:", members)

    // 3. 判断元素是否存在
    exists, err := rdb.SIsMember(ctx, "tags", "redis").Result()
    fmt.Println("SIsMember:", exists)

    // 4. 获取集合元素数量
    count, err := rdb.SCard(ctx, "tags").Result()
    fmt.Println("SCard:", count)

    // 5. 删除元素
    err = rdb.SRem(ctx, "tags", "cache").Err()

    // 6. 随机获取元素
    vals, err := rdb.SRandMemberN(ctx, "tags", 2).Result()
    fmt.Println("SRandMember:", vals)

    // 7. 随机弹出元素
    val, err := rdb.SPop(ctx, "tags").Result()
    fmt.Println("SPop:", val)

    // 8. 移动元素到另一个集合
    success, err := rdb.SMove(ctx, "tags", "newtags", "redis").Result()
    fmt.Println("SMove 成功:", success)

    // 9. 集合交集
    inter, err := rdb.SInter(ctx, "set1", "set2").Result()
    fmt.Println("SInter:", inter)

    // 10. 集合并集
    union, err := rdb.SUnion(ctx, "set1", "set2").Result()
    fmt.Println("SUnion:", union)

    // 11. 集合差集
    diff, err := rdb.SDiff(ctx, "set1", "set2").Result()
    fmt.Println("SDiff:", diff)

    // 12. 将交集结果存储到新集合
    count, err = rdb.SInterStore(ctx, "result", "set1", "set2").Result()
    fmt.Println("SInterStore count:", count)
}
```

## 七、Sorted Set 操作

```go
func ZSetOperations(rdb *redis.Client) {
    // 1. 添加元素
    err := rdb.ZAdd(ctx, "leaderboard", redis.Z{
        Score:  100,
        Member: "player1",
    }, redis.Z{
        Score:  200,
        Member: "player2",
    }, redis.Z{
        Score:  150,
        Member: "player3",
    }).Err()

    // 2. 获取指定范围的元素（升序）
    vals, err := rdb.ZRange(ctx, "leaderboard", 0, -1).Result()
    fmt.Println("ZRange:", vals)

    // 3. 获取指定范围的元素（带分数）
    valsWithScores, err := rdb.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()
    for _, z := range valsWithScores {
        fmt.Printf("Member: %s, Score: %f\n", z.Member, z.Score)
    }

    // 4. 获取指定范围的元素（降序）
    vals, err = rdb.ZRevRange(ctx, "leaderboard", 0, 2).Result()
    fmt.Println("ZRevRange:", vals)

    // 5. 获取指定分数范围的元素
    vals, err = rdb.ZRangeByScore(ctx, "leaderboard", &redis.ZRangeBy{
        Min: "100",
        Max: "200",
    }).Result()
    fmt.Println("ZRangeByScore:", vals)

    // 6. 获取元素的分数
    score, err := rdb.ZScore(ctx, "leaderboard", "player1").Result()
    fmt.Println("ZScore:", score)

    // 7. 获取元素的排名（升序，从0开始）
    rank, err := rdb.ZRank(ctx, "leaderboard", "player1").Result()
    fmt.Println("ZRank:", rank)

    // 8. 获取元素的排名（降序）
    rank, err = rdb.ZRevRank(ctx, "leaderboard", "player1").Result()
    fmt.Println("ZRevRank:", rank)

    // 9. 增加元素的分数
    newScore, err := rdb.ZIncrBy(ctx, "leaderboard", 50, "player1").Result()
    fmt.Println("New Score:", newScore)

    // 10. 获取集合元素数量
    count, err := rdb.ZCard(ctx, "leaderboard").Result()
    fmt.Println("ZCard:", count)

    // 11. 获取指定分数范围的元素数量
    count, err = rdb.ZCount(ctx, "leaderboard", "100", "200").Result()
    fmt.Println("ZCount:", count)

    // 12. 删除元素
    err = rdb.ZRem(ctx, "leaderboard", "player1").Err()

    // 13. 删除指定排名范围的元素
    err = rdb.ZRemRangeByRank(ctx, "leaderboard", 0, 2).Err()

    // 14. 删除指定分数范围的元素
    err = rdb.ZRemRangeByScore(ctx, "leaderboard", "0", "100").Err()
}
```

## 八、通用操作

```go
func CommonOperations(rdb *redis.Client) {
    // 1. 设置过期时间
    err := rdb.Expire(ctx, "key", 10*time.Minute).Err()

    // 2. 查看剩余过期时间
    ttl, err := rdb.TTL(ctx, "key").Result()
    fmt.Println("TTL:", ttl)

    // 3. 移除过期时间
    err = rdb.Persist(ctx, "key").Err()

    // 4. 判断键是否存在
    exists, err := rdb.Exists(ctx, "key").Result()
    fmt.Println("Exists:", exists)

    // 5. 查看键的数据类型
    keyType, err := rdb.Type(ctx, "key").Result()
    fmt.Println("Type:", keyType)

    // 6. 重命名键
    err = rdb.Rename(ctx, "oldkey", "newkey").Err()

    // 7. 删除键
    err = rdb.Del(ctx, "key1", "key2").Err()

    // 8. 查看所有键（慎用）
    keys, err := rdb.Keys(ctx, "*").Result()
    fmt.Println("Keys:", keys)

    // 9. 扫描键（推荐）
    iter := rdb.Scan(ctx, 0, "user:*", 10).Iterator()
    for iter.Next(ctx) {
        fmt.Println("Key:", iter.Val())
    }
    if err := iter.Err(); err != nil {
        panic(err)
    }
}
```


## 九、Pipeline（管道）

Pipeline 可以一次性发送多个命令，减少网络往返时间。

```go
func PipelineExample(rdb *redis.Client) {
    pipe := rdb.Pipeline()

    // 添加多个命令到管道
    incr := pipe.Incr(ctx, "counter")
    pipe.Expire(ctx, "counter", time.Hour)

    // 执行管道
    _, err := pipe.Exec(ctx)
    if err != nil {
        panic(err)
    }

    // 获取命令结果
    fmt.Println("Counter:", incr.Val())
}
```

## 十、事务

```go
func TransactionExample(rdb *redis.Client) {
    // 使用 TxPipeline
    pipe := rdb.TxPipeline()

    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)

    // 执行事务
    _, err := pipe.Exec(ctx)
    if err != nil {
        panic(err)
    }

    // 使用 Watch 实现乐观锁
    err = rdb.Watch(ctx, func(tx *redis.Tx) error {
        // 获取当前值
        val, err := tx.Get(ctx, "balance").Int()
        if err != nil && err != redis.Nil {
            return err
        }

        // 业务逻辑
        val += 100

        // 在事务中更新
        _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
            pipe.Set(ctx, "balance", val, 0)
            return nil
        })
        return err
    }, "balance")

    if err != nil {
        panic(err)
    }
}
```

## 十一、发布订阅

```go
func PubSubExample(rdb *redis.Client) {
    // 订阅频道
    pubsub := rdb.Subscribe(ctx, "news", "sports")
    defer pubsub.Close()

    // 等待订阅确认
    _, err := pubsub.Receive(ctx)
    if err != nil {
        panic(err)
    }

    // 接收消息
    ch := pubsub.Channel()
    go func() {
        for msg := range ch {
            fmt.Printf("Channel: %s, Message: %s\n", msg.Channel, msg.Payload)
        }
    }()

    // 发布消息（在另一个连接中）
    err = rdb.Publish(ctx, "news", "今日新闻").Err()
    if err != nil {
        panic(err)
    }

    time.Sleep(time.Second)
}
```

## 十二、完整示例

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    // 创建 Redis 客户端
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    defer rdb.Close()

    // 测试连接
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("连接成功:", pong)

    // String 操作
    err = rdb.Set(ctx, "user:name", "张三", 10*time.Minute).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "user:name").Result()
    if err == redis.Nil {
        fmt.Println("键不存在")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("user:name =", val)
    }

    // Hash 操作
    err = rdb.HSet(ctx, "user:1", map[string]interface{}{
        "name":  "李四",
        "age":   30,
        "email": "lisi@example.com",
    }).Err()
    if err != nil {
        panic(err)
    }

    user, err := rdb.HGetAll(ctx, "user:1").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("用户信息:", user)

    // List 操作
    err = rdb.RPush(ctx, "tasks", "task1", "task2", "task3").Err()
    if err != nil {
        panic(err)
    }

    tasks, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("任务列表:", tasks)

    // Set 操作
    err = rdb.SAdd(ctx, "tags", "go", "redis", "database").Err()
    if err != nil {
        panic(err)
    }

    tags, err := rdb.SMembers(ctx, "tags").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("标签:", tags)

    // Sorted Set 操作
    err = rdb.ZAdd(ctx, "scores", redis.Z{Score: 90, Member: "张三"}).Err()
    err = rdb.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "李四"}).Err()
    err = rdb.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "王五"}).Err()

    scores, err := rdb.ZRevRangeWithScores(ctx, "scores", 0, -1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("成绩排名:")
    for i, z := range scores {
        fmt.Printf("%d. %s: %.0f\n", i+1, z.Member, z.Score)
    }
}
```

## 十三、错误处理

```go
import "github.com/redis/go-redis/v9"

// 判断键是否不存在
val, err := rdb.Get(ctx, "key").Result()
if err == redis.Nil {
    fmt.Println("键不存在")
} else if err != nil {
    panic(err)
} else {
    fmt.Println("key =", val)
}
```

## 十四、最佳实践

1. 使用连接池，避免频繁创建连接
2. 使用 Pipeline 批量执行命令
3. 合理设置过期时间，避免内存溢出
4. 使用 Scan 代替 Keys 遍历键
5. 错误处理要区分 redis.Nil 和其他错误
6. 使用 context 控制超时
7. 生产环境建议使用 Redis Cluster 或哨兵模式
