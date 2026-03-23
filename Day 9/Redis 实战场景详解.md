# Redis 实战场景详解

## 先理解 Redis 的定位

```
┌─────────────────────────────────────────────────────┐
│                    你的应用架构                        │
│                                                       │
│   用户请求 ──→ [ 应用服务器 ] ──→ [ MySQL/PostgreSQL ]  │
│                    │                    ↑              │
│                    │                    │              │
│                    ▼                    │              │
│              [ Redis ] ◄── 热数据放这里，快100倍        │
│              (内存, <1ms)                              │
└─────────────────────────────────────────────────────┘
```

**核心思想**：MySQL 磁盘IO慢（毫秒级），Redis 内存操作快（微秒级）。把高频访问的数据放 Redis，减轻数据库压力。

---

## 场景一：缓存穿透防护 — 电商商品详情页

**问题**：商品详情页每秒被访问上万次，每次都查 MySQL 会崩。

```
用户请求商品详情
      │
      ▼
  查 Redis 有缓存吗？
     ╱  ╲
   有     没有
   │       │
   ▼       ▼
 直接返回  查 MySQL → 写入 Redis → 返回
```

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

// 模拟的商品结构
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Stock int     `json:"stock"`
}

// ============ 方案一：最基础的缓存 (Cache-Aside) ============
func GetProductBasic(productID int) (*Product, error) {
    cacheKey := fmt.Sprintf("product:%d", productID)

    // 第一步：先查 Redis
    val, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        // 缓存命中，直接返回
        var p Product
        json.Unmarshal([]byte(val), &p)
        fmt.Println(">>> 缓存命中，来自 Redis")
        return &p, nil
    }

    // 第二步：缓存没命中，查数据库（这里模拟）
    fmt.Println(">>> 缓存未命中，查数据库")
    product := queryProductFromDB(productID)

    // 第三步：写入缓存，设置过期时间（防止数据永远不更新）
    data, _ := json.Marshal(product)
    rdb.Set(ctx, cacheKey, data, 10*time.Minute) // 10分钟过期

    return product, nil
}

// ============ 方案二：防止缓存穿透（空值缓存） ============
// 缓存穿透：恶意请求不存在的ID，每次绕过缓存直接打数据库
func GetProductWithPenetrationGuard(productID int) (*Product, error) {
    cacheKey := fmt.Sprintf("product:%d", productID)

    val, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        // 特殊处理：空值标记
        if val == "NULL" {
            fmt.Println(">>> 命中空值缓存，商品不存在")
            return nil, fmt.Errorf("product not found")
        }
        var p Product
        json.Unmarshal([]byte(val), &p)
        return &p, nil
    }

    product := queryProductFromDB(productID)
    if product == nil {
        // 商品不存在，缓存一个空值，过期时间短一些
        rdb.Set(ctx, cacheKey, "NULL", 5*time.Minute)
        return nil, fmt.Errorf("product not found")
    }

    data, _ := json.Marshal(product)
    rdb.Set(ctx, cacheKey, data, 10*time.Minute)
    return product, nil
}

// ============ 方案三：防止缓存雪崩（随机过期时间） ============
// 缓存雪崩：大量key同时过期，瞬间全部打到数据库
func GetProductWithAvalancheGuard(productID int) (*Product, error) {
    cacheKey := fmt.Sprintf("product:%d", productID)

    val, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var p Product
        json.Unmarshal([]byte(val), &p)
        return &p, nil
    }

    product := queryProductFromDB(productID)

    // 关键：基础过期时间 + 随机偏移，避免同时过期
    baseTTL := 10 * time.Minute
    jitter := time.Duration(productID%120) * time.Second // 0~120秒随机偏移
    rdb.Set(ctx, cacheKey, mustMarshal(product), baseTTL+jitter)

    return product, nil
}

// ============ 方案四：缓存击穿防护（分布式锁） ============
// 缓存击穿：某个热点key过期的瞬间，大量请求同时涌入
func GetProductWithHotKeyGuard(productID int) (*Product, error) {
    cacheKey := fmt.Sprintf("product:%d", productID)
    lockKey := fmt.Sprintf("lock:product:%d", productID)

    // 查缓存
    val, err := rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        var p Product
        json.Unmarshal([]byte(val), &p)
        return &p, nil
    }

    // 缓存没命中，尝试获取分布式锁
    locked, _ := rdb.SetNX(ctx, lockKey, "1", 10*time.Second).Result()
    if locked {
        // 我抢到了锁，我去查数据库并更新缓存
        fmt.Println(">>> 我抢到锁了，我去查DB")
        product := queryProductFromDB(productID)
        data, _ := json.Marshal(product)
        rdb.Set(ctx, cacheKey, data, 10*time.Minute)
        rdb.Del(ctx, lockKey) // 释放锁
        return product, nil
    }

    // 没抢到锁，等一会儿重试
    fmt.Println(">>> 没抢到锁，等一下再试")
    time.Sleep(100 * time.Millisecond)
    return GetProductWithHotKeyGuard(productID) // 重试
}

func queryProductFromDB(id int) *Product {
    // 模拟数据库查询
    time.Sleep(50 * time.Millisecond) // 模拟数据库延迟
    if id == 404 {
        return nil
    }
    return &Product{ID: id, Name: fmt.Sprintf("商品%d", id), Price: 99.9, Stock: 100}
}

func mustMarshal(v interface{}) []byte {
    data, _ := json.Marshal(v)
    return data
}
```

---

## 场景二：分布式 Session — 多服务器登录状态共享

**问题**：你的应用有3台服务器，用户登录在服务器A，下次请求到了服务器B，B不认识这个用户。

```
以前（单机）：Session 存在服务器内存里
现在（多机）：Session 存在 Redis 里，所有服务器共享
```

```
用户登录服务器A → 生成 session_id → 存入 Redis
用户请求服务器B → 带上 session_id → 从 Redis 读取 → 认识这个用户
```

```go
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// 用户登录：生成 session 存入 Redis
func Login(userID int, username string) (string, error) {
    // 生成随机 session ID
    sessionID := generateSessionID()

    // 把用户信息存入 Redis，key 是 sessionID
    sessionKey := fmt.Sprintf("session:%s", sessionID)
    sessionData := fmt.Sprintf(`{"user_id":%d,"username":"%s","login_time":"%s"}`,
        userID, username, time.Now().Format(time.RFC3339))

    // Session 30分钟过期（滑动过期）
    err := rdb.Set(ctx, sessionKey, sessionData, 30*time.Minute).Err()
    if err != nil {
        return "", err
    }

    return sessionID, nil
}

// 验证 Session：任何服务器都能验证
func ValidateSession(sessionID string) (map[string]interface{}, error) {
    sessionKey := fmt.Sprintf("session:%s", sessionID)

    val, err := rdb.Get(ctx, sessionKey).Result()
    if err != nil {
        return nil, fmt.Errorf("session expired or invalid")
    }

    // 每次访问都续期（滑动过期）
    rdb.Expire(ctx, sessionKey, 30*time.Minute)

    fmt.Printf(">>> Session 有效: %s\n", val)
    return nil, nil
}

// 退出登录
func Logout(sessionID string) {
    sessionKey := fmt.Sprintf("session:%s", sessionID)
    rdb.Del(ctx, sessionKey)
}

// 获取所有在线用户（管理后台用）
func GetOnlineUserCount() int64 {
    // 用 SCAN 而不是 KEYS，避免阻塞 Redis
    var cursor uint64
    var count int64
    for {
        keys, nextCursor, _ := rdb.Scan(ctx, cursor, "session:*", 100).Result()
        count += int64(len(keys))
        cursor = nextCursor
        if cursor == 0 {
            break
        }
    }
    return count
}

func generateSessionID() string {
    bytes := make([]byte, 32)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}
```

---

## 场景三：API 限流 — 保护你的接口不被打爆

**问题**：你的API对外开放，有人恶意刷接口，或者某个用户请求太频繁。

```go
package main

import (
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// ============ 方案一：固定窗口计数器 ============
// 限制每个用户每分钟最多 100 次请求
func RateLimitFixedWindow(userID string) bool {
    key := fmt.Sprintf("ratelimit:%s:%s", userID, time.Now().Format("200601021504"))

    // 自增计数
    count, _ := rdb.Incr(ctx, key).Result()

    // 第一次请求时设置过期时间
    if count == 1 {
        rdb.Expire(ctx, key, time.Minute)
    }

    if count > 100 {
        fmt.Printf("用户 %s 被限流了！当前次数: %d\n", userID, count)
        return false // 拒绝请求
    }

    fmt.Printf("用户 %s 请求通过，当前次数: %d\n", userID, count)
    return true
}

// ============ 方案二：滑动窗口（更精确） ============
// 用 Sorted Set 实现，记录每次请求的时间戳
func RateLimitSlidingWindow(userID string) bool {
    key := fmt.Sprintf("ratelimit:sliding:%s", userID)
    now := float64(time.Now().UnixMilli())
    windowStart := now - 60000 // 60秒窗口

    // 清理窗口外的旧记录
    rdb.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%f", windowStart))

    // 获取当前窗口内的请求数
    count, _ := rdb.ZCard(ctx, key).Result()

    if count >= 100 {
        fmt.Printf("用户 %s 被限流（滑动窗口）！\n", userID)
        return false
    }

    // 记录本次请求
    rdb.ZAdd(ctx, key, redis.Z{Score: now, Member: fmt.Sprintf("%f", now)})
    rdb.Expire(ctx, key, time.Minute) // 防止key永远不删除

    return true
}

// ============ 方案三：令牌桶算法（允许突发流量） ============
// 每秒补充10个令牌，桶容量100
func RateLimitTokenBucket(userID string) bool {
    key := fmt.Sprintf("bucket:%s", userID)

    // 用 Lua 脚本保证原子性
    luaScript := redis.NewScript(`
        local key = KEYS[1]
        local rate = tonumber(ARGV[1])       -- 每秒补充速率
        local capacity = tonumber(ARGV[2])   -- 桶容量
        local now = tonumber(ARGV[3])        -- 当前时间戳(秒)
        local requested = tonumber(ARGV[4])  -- 本次消耗令牌数

        local data = redis.call('HMGET', key, 'tokens', 'last_time')
        local tokens = tonumber(data[1]) or capacity
        local last_time = tonumber(data[2]) or now

        -- 计算从上次到现在补充了多少令牌
        local elapsed = now - last_time
        tokens = math.min(capacity, tokens + elapsed * rate)

        if tokens >= requested then
            tokens = tokens - requested
            redis.call('HMSET', key, 'tokens', tokens, 'last_time', now)
            redis.call('EXPIRE', key, 60)
            return 1  -- 允许
        else
            redis.call('HMSET', key, 'tokens', tokens, 'last_time', now)
            redis.call('EXPIRE', key, 60)
            return 0  -- 拒绝
        end
    `)

    now := float64(time.Now().Unix())
    result, _ := luaScript.Run(ctx, rdb, []string{key}, 10, 100, now, 1).Int()

    return result == 1
}
```

---

## 场景四：排行榜 — 游戏积分 / 热搜榜

**这是 Sorted Set 最经典的场景。**

```go
package main

import (
    "fmt"

    "github.com/redis/go-redis/v9"
)

// ============ 实时排行榜 ============
func GameLeaderboard() {
    key := "game:leaderboard"

    // 1. 玩家得分更新（玩一局加分数）
    updateScore(key, "player:张三", 150)
    updateScore(key, "player:李四", 320)
    updateScore(key, "player:王五", 280)
    updateScore(key, "player:赵六", 410)
    updateScore(key, "player:孙七", 190)

    // 2. 获取 Top 10（从高到低）
    top10, _ := rdb.ZRevRangeWithScores(ctx, key, 0, 9).Result()
    fmt.Println("===== 排行榜 Top 10 =====")
    for i, z := range top10 {
        fmt.Printf("第%d名: %s  %.0f分\n", i+1, z.Member, z.Score)
    }

    // 3. 查看我的排名
    myRank, _ := rdb.ZRevRank(ctx, key, "player:张三").Result()
    myScore, _ := rdb.ZScore(ctx, key, "player:张三").Result()
    fmt.Printf("\n张三的排名: 第%d名, 分数: %.0f\n", myRank+1, myScore)

    // 4. 查看我周围的人（上下各2名）
    rank, _ := rdb.ZRevRank(ctx, key, "player:张三").Result()
    start := rank - 2
    end := rank + 2
    if start < 0 {
        start = 0
    }
    around, _ := rdb.ZRevRangeWithScores(ctx, key, start, end).Result()
    fmt.Println("\n===== 张三周围的玩家 =====")
    for _, z := range around {
        fmt.Printf("  %s  %.0f分\n", z.Member, z.Score)
    }

    // 5. 按分数段查询（比如 200~400分的玩家）
    midPlayers, _ := rdb.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
        Min: "200",
        Max: "400",
    }).Result()
    fmt.Println("\n===== 200~400分的玩家 =====")
    for _, z := range midPlayers {
        fmt.Printf("  %s  %.0f分\n", z.Member, z.Score)
    }
}

func updateScore(key, member string, delta float64) {
    // ZINCRBY：分数累加（玩一局加一次分）
    rdb.ZIncrBy(ctx, key, delta, member)
}
```

---

## 场景五：延迟任务队列 — 订单30分钟未支付自动取消

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// ============ 用 Sorted Set 实现延迟队列 ============
// score = 应该被处理的时间戳

type DelayTask struct {
    TaskID   string `json:"task_id"`
    TaskType string `json:"task_type"` // "cancel_order", "send_notification"...
    Payload  string `json:"payload"`
}

// 生产者：添加延迟任务
func AddDelayTask(task DelayTask, delay time.Duration) {
    key := "delay_queue"
    executeAt := float64(time.Now().Add(delay).Unix())

    data, _ := json.Marshal(task)
    rdb.ZAdd(ctx, key, redis.Z{
        Score:  executeAt,
        Member: string(data),
    })

    fmt.Printf("任务 %s 已加入队列，将在 %v 后执行\n", task.TaskID, delay)
}

// 消费者：轮询处理到期任务
func ProcessDelayTasks() {
    key := "delay_queue"

    for {
        now := float64(time.Now().Unix())

        // 取出所有已到期的任务（score <= 当前时间），最多取10条
        tasks, _ := rdb.ZRangeByScore(ctx, key, &redis.ZRangeBy{
            Min:    "-inf",
            Max:    fmt.Sprintf("%f", now),
            Offset: 0,
            Count:  10,
        }).Result()

        if len(tasks) == 0 {
            time.Sleep(1 * time.Second) // 没有任务，休息1秒
            continue
        }

        for _, taskStr := range tasks {
            var task DelayTask
            json.Unmarshal([]byte(taskStr), &task)

            // 原子性地删除（防止重复消费）
            removed, _ := rdb.ZRem(ctx, key, taskStr).Result()
            if removed == 0 {
                continue // 别的消费者已经处理了
            }

            // 执行任务
            executeTask(task)
        }
    }
}

func executeTask(task DelayTask) {
    switch task.TaskType {
    case "cancel_order":
        fmt.Printf(">>> 取消订单: %s (payload: %s)\n", task.TaskID, task.Payload)
        // 实际业务：更新订单状态为"已取消"，恢复库存...
    case "send_reminder":
        fmt.Printf(">>> 发送提醒: %s\n", task.TaskID)
    }
}

// 使用示例
func DelayQueueDemo() {
    // 用户下单，30分钟后未支付则取消
    AddDelayTask(DelayTask{
        TaskID:   "order_1001",
        TaskType: "cancel_order",
        Payload:  `{"order_id":1001,"user_id":42}`,
    }, 30*time.Minute)

    // 5秒后发提醒（演示用）
    AddDelayTask(DelayTask{
        TaskID:   "reminder_1001",
        TaskType: "send_reminder",
        Payload:  `{"msg":"您的订单即将超时"}`,
    }, 5*time.Second)

    // 启动消费者
    go ProcessDelayTasks()
}
```

---

## 场景六：Feed 流 / 朋友圈 — 关注人的动态

```go
package main

import (
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// ============ 推模式：发微博时推送到粉丝的 Timeline ============

// 用户发微博
func PostWeibo(userID int, content string) {
    // 1. 生成微博ID（实际用数据库自增ID或雪花算法）
    postID := time.Now().UnixNano()

    // 2. 存储微博内容
    rdb.HSet(ctx, fmt.Sprintf("post:%d", postID),
        "user_id", userID,
        "content", content,
        "time", time.Now().Format(time.RFC3339),
    )

    // 3. 把微博ID推送到所有粉丝的 Timeline（关键步骤）
    fanKey := fmt.Sprintf("fans:%d", userID)
    fans, _ := rdb.SMembers(ctx, fanKey).Result()

    for _, fanID := range fans {
        timelineKey := fmt.Sprintf("timeline:%s", fanID)
        // LPUSH 推到最前面（最新的在前面）
        rdb.LPush(ctx, timelineKey, postID)
        // 只保留最近 500 条
        rdb.LTrim(ctx, timelineKey, 0, 499)
    }

    fmt.Printf("用户%d发了微博，推送给%d个粉丝\n", userID, len(fans))
}

// 用户刷朋友圈/Timeline
func GetTimeline(userID int, page, pageSize int) {
    timelineKey := fmt.Sprintf("timeline:%d", userID)

    start := int64((page - 1) * pageSize)
    end := start + int64(pageSize) - 1

    // 从 Timeline 中取微博ID
    postIDs, _ := rdb.LRange(ctx, timelineKey, start, end).Result()

    fmt.Printf("===== 用户%d的 Timeline (第%d页) =====\n", userID, page)
    for _, postID := range postIDs {
        // 根据ID获取微博详情
        post, _ := rdb.HGetAll(ctx, fmt.Sprintf("post:%s", postID)).Result()
        fmt.Printf("  [%s] 用户%s: %s\n",
            post["time"], post["user_id"], post["content"])
    }
}

// 关注某人
func Follow(fanID, targetID int) {
    // 我的关注列表
    rdb.SAdd(ctx, fmt.Sprintf("following:%d", fanID), targetID)
    // 对方的粉丝列表
    rdb.SAdd(ctx, fmt.Sprintf("fans:%d", targetID), fanID)
}
```

---

## 场景七：分布式锁 — 秒杀/库存扣减

**问题**：100件商品，1000人同时抢，不能超卖。

```go
package main

import (
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// ============ 简单版：用 SETNX 实现分布式锁 ============
func SeckillSimple(productID int, userID int) bool {
    lockKey := fmt.Sprintf("lock:seckill:%d", productID)
    stockKey := fmt.Sprintf("stock:%d", productID)

    // 1. 尝试加锁
    locked, _ := rdb.SetNX(ctx, lockKey, userID, 10*time.Second).Result()
    if !locked {
        fmt.Printf("用户%d: 没抢到锁\n", userID)
        return false
    }
    defer rdb.Del(ctx, lockKey) // 用 defer 确保释放锁

    // 2. 查看库存
    stock, _ := rdb.Get(ctx, stockKey).Int()
    if stock <= 0 {
        fmt.Printf("用户%d: 库存不足\n", userID)
        return false
    }

    // 3. 扣减库存
    rdb.Decr(ctx, stockKey)
    fmt.Printf("用户%d: 抢购成功！剩余库存: %d\n", userID, stock-1)
    return true
}

// ============ 进阶版：用 Lua 脚本保证原子性 ============
// Lua 脚本在 Redis 中是原子执行的，不会被打断
func SeckillAtomic(productID int, userID int) bool {
    luaScript := redis.NewScript(`
        local stockKey = KEYS[1]
        local lockKey = KEYS[2]
        
        -- 检查库存
        local stock = tonumber(redis.call('GET', stockKey) or '0')
        if stock <= 0 then
            return 0  -- 库存不足
        end
        
        -- 检查是否已抢过（一人一单）
        local bought = redis.call('SISMEMBER', lockKey, ARGV[1])
        if bought == 1 then
            return -1  -- 已抢过
        end
        
        -- 扣库存 + 记录用户
        redis.call('DECR', stockKey)
        redis.call('SADD', lockKey, ARGV[1])
        return 1  -- 成功
    `)

    stockKey := fmt.Sprintf("stock:%d", productID)
    boughtKey := fmt.Sprintf("seckill:users:%d", productID)

    result, _ := luaScript.Run(ctx, rdb, []string{stockKey, boughtKey}, userID).Int()

    switch result {
    case 1:
        fmt.Printf("用户%d: 秒杀成功！\n", userID)
        return true
    case 0:
        fmt.Printf("用户%d: 库存已空\n", userID)
        return false
    case -1:
        fmt.Printf("用户%d: 你已经抢过了\n", userID)
        return false
    }
    return false
}
```

---

## 场景八：UV/PV 统计 — 网站实时访问量

```go
package main

import (
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

// ============ PV (页面访问量)：用 String + INCR 就行 ============
func RecordPV(page string) {
    today := time.Now().Format("20060102")
    key := fmt.Sprintf("pv:%s:%s", page, today)
    rdb.Incr(ctx, key)
    rdb.Expire(ctx, key, 72*time.Hour) // 保留3天
}

// ============ UV (独立访客)：用 HyperLogLog，省内存 ============
// 1000万用户只占 12KB 内存！
func RecordUV(page string, userID string) {
    today := time.Now().Format("20060102")
    key := fmt.Sprintf("uv:%s:%s", page, today)
    rdb.PFAdd(ctx, key, userID)
    rdb.Expire(ctx, key, 72*time.Hour)
}

func GetUV(page string) int64 {
    today := time.Now().Format("20060102")
    key := fmt.Sprintf("uv:%s:%s", page, today)
    count, _ := rdb.PFCount(ctx, key).Result()
    return count
}

// ============ 用户签到：用 Bitmap ============
func UserSign(userID int) {
    month := time.Now().Format("200601")
    key := fmt.Sprintf("sign:%d:%s", userID, month)
    day := time.Now().Day() - 1 // Bitmap 从0开始

    rdb.SetBit(ctx, key, int64(day), 1)
    rdb.Expire(ctx, key, 90*24*time.Hour)
}

// 查看本月签到天数
func GetSignCount(userID int) int64 {
    month := time.Now().Format("200601")
    key := fmt.Sprintf("sign:%d:%s", userID, month)
    count, _ := rdb.BitCount(ctx, key, nil).Result()
    return count
}

// 查看今天是否签到
func IsSignedToday(userID int) bool {
    month := time.Now().Format("200601")
    key := fmt.Sprintf("sign:%d:%s", userID, month)
    day := time.Now().Day() - 1

    bit, _ := rdb.GetBit(ctx, key, int64(day)).Result()
    return bit == 1
}
```

---

## 场景九：布隆过滤器 — 判断元素是否"可能存在"

**问题**：10亿个已知恶意IP，快速判断一个IP是否在黑名单中。用 Set 存10亿个太占内存。

```go
// Redis 4.0+ 支持布隆过滤器模块
// 或者用 go-redis 客户端模拟

func BloomFilterDemo() {
    // 使用 RedisBloom 模块（需要安装）
    // BF.ADD blacklist_ip 192.168.1.1
    // BF.EXISTS blacklist_ip 192.168.1.1

    // 如果没有模块，可以用多个 Hash + 多个哈希函数模拟
    // 或者用 SETBIT 在 Bitmap 上模拟简易布隆过滤器
}
```

---

## 全景总结：一张图看懂 Redis 在系统中的位置

```
                         ┌─────────────────────────┐
                         │       用户请求            │
                         └───────────┬─────────────┘
                                     │
                         ┌───────────▼─────────────┐
                         │       Nginx 网关         │
                         │  (IP限流 → Redis计数)     │
                         └───────────┬─────────────┘
                                     │
                    ┌────────────────┼────────────────┐
                    ▼                ▼                 ▼
             ┌────────────┐  ┌────────────┐   ┌────────────┐
             │  服务器 A   │  │  服务器 B   │   │  服务器 C   │
             └─────┬──────┘  └─────┬──────┘   └─────┬──────┘
                   │               │                 │
                   └───────────────┼─────────────────┘
                                   │
                    ┌──────────────▼──────────────┐
                    │          Redis 集群          │
                    │                             │
                    │  ┌───────────────────────┐  │
                    │  │ String: 缓存/计数器/锁  │  │
                    │  │ List:   队列/Timeline  │  │
                    │  │ Hash:   Session/对象    │  │
                    │  │ Set:    标签/共同关注    │  │
                    │  │ ZSet:   排行榜/延迟队列  │  │
                    │  └───────────────────────┘  │
                    └──────────────┬──────────────┘
                                   │ (缓存没有才查这里)
                    ┌──────────────▼──────────────┐
                    │      MySQL / PostgreSQL      │
                    │      (持久化存储，最终数据源)  │
                    └─────────────────────────────┘
```

| 场景 | Redis 数据类型 | 为什么用 Redis |
|:---|:---|:---|
| 商品缓存 | String | 内存读取 <1ms vs 数据库 50ms+ |
| 用户Session | String/Hash | 多服务器共享登录状态 |
| API限流 | String(INCR) / ZSet | 原子计数，精确控制频率 |
| 排行榜 | Sorted Set | 天然支持排序和范围查询 |
| 延迟任务 | Sorted Set | 按时间戳排序，到期取出 |
| 朋友圈/Feed | List | 有序，支持分页 |
| 秒杀库存 | String(INCR/DECR) + Lua | 原子操作，不超卖 |
| UV统计 | HyperLogLog | 12KB统计千万级去重 |
| 签到打卡 | Bitmap | 一个用户一年只需46字节 |
| 黑名单判断 | 布隆过滤器 | 1亿数据只需100MB |

**记住一句话**：**Redis 不是用来替代数据库的，而是数据库前面的"高速缓冲层"和"多功能工具箱"**。