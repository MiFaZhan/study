package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StringOperations(rdb *redis.Client) {
	fmt.Println("\n========== String 操作 ==========")

	// 1. 设置和获取值
	rdb.Set(ctx, "name", "张三", 0)
	val, _ := rdb.Get(ctx, "name").Result()
	fmt.Printf("  GET name: %s\n", val)

	// 2. SetNX (仅当 key 不存在时设置)
	success, _ := rdb.SetNX(ctx, "lock", "locked", 30*time.Second).Result()
	fmt.Printf("  SETNX lock: %v\n", success)

	// 3. 批量操作
	rdb.MSet(ctx, "key1", "value1", "key2", "value2")
	vals, _ := rdb.MGet(ctx, "key1", "key2").Result()
	fmt.Printf("  MGET key1, key2: %v\n", vals)

	// 4. 递增操作
	count, _ := rdb.Incr(ctx, "counter").Result()
	fmt.Printf("  INCR counter: %d\n", count)
	count, _ = rdb.IncrBy(ctx, "counter", 10).Result()
	fmt.Printf("  INCRBY counter 10: %d\n", count)

	// 5. 字符串追加
	length, _ := rdb.Append(ctx, "name", "先生").Result()
	fmt.Printf("  APPEND name '先生': 长度=%d\n", length)

	// 清理
	rdb.Del(ctx, "key1", "key2", "counter", "lock", "session")
}
