# Redis 各种类型操作详解

## 一、Redis 简介

Redis（Remote Dictionary Server）是一个开源的内存数据结构存储系统，可以用作数据库、缓存和消息代理。

## 二、Redis 数据类型及操作

### 1. String（字符串）

String 是 Redis 最基本的数据类型，一个 key 对应一个 value。

#### 基本操作

```bash
# 设置值
SET key value
SET name "张三"

# 获取值
GET key
GET name

# 设置多个键值对
MSET key1 value1 key2 value2
MSET name "张三" age "25"

# 获取多个值
MGET key1 key2
MGET name age

# 设置值并设置过期时间（秒）
SETEX key seconds value
SETEX session:token 3600 "abc123"

# 仅当 key 不存在时设置
SETNX key value
SETNX lock:user:1 "locked"

# 追加字符串
APPEND key value
APPEND name "先生"

# 获取字符串长度
STRLEN key
STRLEN name

# 删除键
DEL key
DEL name
```

#### 数值操作

```bash
# 递增
INCR key
INCR counter

# 递增指定值
INCRBY key increment
INCRBY counter 10

# 递减
DECR key
DECR counter

# 递减指定值
DECRBY key decrement
DECRBY counter 5

# 浮点数增加
INCRBYFLOAT key increment
INCRBYFLOAT price 0.5
```

### 2. Hash（哈希）

Hash 是一个 string 类型的 field 和 value 的映射表，适合存储对象。

```bash
# 设置单个字段
HSET key field value
HSET user:1 name "张三"
HSET user:1 age 25

# 获取单个字段
HGET key field
HGET user:1 name

# 设置多个字段
HMSET key field1 value1 field2 value2
HMSET user:2 name "李四" age 30 email "lisi@example.com"

# 获取多个字段
HMGET key field1 field2
HMGET user:2 name age

# 获取所有字段和值
HGETALL key
HGETALL user:1

# 获取所有字段名
HKEYS key
HKEYS user:1

# 获取所有值
HVALS key
HVALS user:1

# 判断字段是否存在
HEXISTS key field
HEXISTS user:1 name

# 删除字段
HDEL key field1 field2
HDEL user:1 email

# 获取字段数量
HLEN key
HLEN user:1

# 字段值递增
HINCRBY key field increment
HINCRBY user:1 age 1

# 仅当字段不存在时设置
HSETNX key field value
HSETNX user:1 phone "13800138000"
```

### 3. List（列表）

List 是简单的字符串列表，按照插入顺序排序。

```bash
# 从左侧插入
LPUSH key value1 value2
LPUSH tasks "task1" "task2"

# 从右侧插入
RPUSH key value1 value2
RPUSH tasks "task3" "task4"

# 从左侧弹出
LPOP key
LPOP tasks

# 从右侧弹出
RPOP key
RPOP tasks

# 获取列表长度
LLEN key
LLEN tasks

# 获取指定范围的元素（0 到 -1 表示所有）
LRANGE key start stop
LRANGE tasks 0 -1
LRANGE tasks 0 2

# 获取指定索引的元素
LINDEX key index
LINDEX tasks 0

# 设置指定索引的值
LSET key index value
LSET tasks 0 "new_task"

# 在指定元素前/后插入
LINSERT key BEFORE|AFTER pivot value
LINSERT tasks BEFORE "task2" "task1.5"

# 删除指定数量的元素
LREM key count value
LREM tasks 1 "task1"  # 删除1个"task1"

# 保留指定范围的元素
LTRIM key start stop
LTRIM tasks 0 9  # 只保留前10个元素

# 阻塞式弹出（用于队列）
BLPOP key timeout
BLPOP tasks 10  # 等待10秒

BRPOP key timeout
BRPOP tasks 10
```

### 4. Set（集合）

Set 是 string 类型的无序集合，不允许重复元素。

```bash
# 添加元素
SADD key member1 member2
SADD tags "redis" "database" "cache"

# 获取所有元素
SMEMBERS key
SMEMBERS tags

# 判断元素是否存在
SISMEMBER key member
SISMEMBER tags "redis"

# 获取集合元素数量
SCARD key
SCARD tags

# 删除元素
SREM key member1 member2
SREM tags "cache"

# 随机获取元素
SRANDMEMBER key count
SRANDMEMBER tags 2

# 随机弹出元素
SPOP key count
SPOP tags 1

# 移动元素到另一个集合
SMOVE source destination member
SMOVE tags newtags "redis"

# 集合运算 - 交集
SINTER key1 key2
SINTER set1 set2

# 集合运算 - 并集
SUNION key1 key2
SUNION set1 set2

# 集合运算 - 差集
SDIFF key1 key2
SDIFF set1 set2

# 将交集结果存储到新集合
SINTERSTORE destination key1 key2
SINTERSTORE result set1 set2
```

### 5. Sorted Set（有序集合）

Sorted Set 是 Set 的升级版，每个元素关联一个分数（score），按分数排序。

```bash
# 添加元素
ZADD key score1 member1 score2 member2
ZADD leaderboard 100 "player1" 200 "player2" 150 "player3"

# 获取指定范围的元素（按分数升序）
ZRANGE key start stop [WITHSCORES]
ZRANGE leaderboard 0 -1
ZRANGE leaderboard 0 -1 WITHSCORES

# 获取指定范围的元素（按分数降序）
ZREVRANGE key start stop [WITHSCORES]
ZREVRANGE leaderboard 0 2 WITHSCORES

# 获取指定分数范围的元素
ZRANGEBYSCORE key min max [WITHSCORES]
ZRANGEBYSCORE leaderboard 100 200 WITHSCORES

# 获取元素的分数
ZSCORE key member
ZSCORE leaderboard "player1"

# 获取元素的排名（升序，从0开始）
ZRANK key member
ZRANK leaderboard "player1"

# 获取元素的排名（降序）
ZREVRANK key member
ZREVRANK leaderboard "player1"

# 增加元素的分数
ZINCRBY key increment member
ZINCRBY leaderboard 50 "player1"

# 获取集合元素数量
ZCARD key
ZCARD leaderboard

# 获取指定分数范围的元素数量
ZCOUNT key min max
ZCOUNT leaderboard 100 200

# 删除元素
ZREM key member1 member2
ZREM leaderboard "player1"

# 删除指定排名范围的元素
ZREMRANGEBYRANK key start stop
ZREMRANGEBYRANK leaderboard 0 2

# 删除指定分数范围的元素
ZREMRANGEBYSCORE key min max
ZREMRANGEBYSCORE leaderboard 0 100
```

## 三、Redis 通用命令

```bash
# 查看所有键
KEYS pattern
KEYS *
KEYS user:*

# 判断键是否存在
EXISTS key
EXISTS name

# 设置过期时间（秒）
EXPIRE key seconds
EXPIRE session:token 3600

# 设置过期时间（毫秒）
PEXPIRE key milliseconds
PEXPIRE session:token 3600000

# 查看剩余过期时间（秒）
TTL key
TTL session:token

# 查看剩余过期时间（毫秒）
PTTL key
PTTL session:token

# 移除过期时间
PERSIST key
PERSIST session:token

# 重命名键
RENAME key newkey
RENAME oldname newname

# 仅当新键不存在时重命名
RENAMENX key newkey
RENAMENX oldname newname

# 查看键的数据类型
TYPE key
TYPE name

# 删除键
DEL key1 key2
DEL name age

# 随机返回一个键
RANDOMKEY

# 切换数据库（0-15）
SELECT index
SELECT 1

# 移动键到其他数据库
MOVE key db
MOVE name 1

# 清空当前数据库
FLUSHDB

# 清空所有数据库
FLUSHALL

# 查看数据库键数量
DBSIZE
```

## 四、Redis 事务

```bash
# 开启事务
MULTI

# 执行命令（加入队列）
SET key1 value1
SET key2 value2

# 执行事务
EXEC

# 取消事务
DISCARD

# 监视键（乐观锁）
WATCH key1 key2

# 取消监视
UNWATCH
```

## 五、Redis 发布订阅

```bash
# 订阅频道
SUBSCRIBE channel1 channel2
SUBSCRIBE news sports

# 发布消息
PUBLISH channel message
PUBLISH news "今日新闻"

# 订阅模式匹配的频道
PSUBSCRIBE pattern
PSUBSCRIBE news:*

# 取消订阅
UNSUBSCRIBE channel1 channel2

# 取消模式订阅
PUNSUBSCRIBE pattern
```

## 六、Redis 持久化

### RDB（快照）

```bash
# 手动触发快照
SAVE      # 阻塞式保存
BGSAVE    # 后台保存

# 查看最后一次保存时间
LASTSAVE
```

### AOF（追加文件）

配置文件设置：
```
appendonly yes
appendfsync everysec  # always/everysec/no
```

## 七、Redis 性能测试与监控

```bash
# 测试连接
PING

# 查看服务器信息
INFO
INFO server
INFO stats
INFO memory

# 实时监控命令
MONITOR

# 查看慢查询日志
SLOWLOG GET 10

# 查看客户端连接
CLIENT LIST

# 关闭客户端连接
CLIENT KILL ip:port
```

## 八、使用场景

1. **String**: 缓存、计数器、分布式锁
2. **Hash**: 存储对象（用户信息、商品信息）
3. **List**: 消息队列、最新列表、排行榜
4. **Set**: 标签、共同好友、去重
5. **Sorted Set**: 排行榜、延时队列、范围查询
