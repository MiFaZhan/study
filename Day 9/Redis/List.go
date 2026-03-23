package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func ListOperations(rdb *redis.Client) {
	fmt.Println("\n========== List 操作 ==========")

	key := "list"
	rdb.Del(ctx, key)

	// 1. LPUSH & RPUSH
	rdb.LPush(ctx, key, "item1", "item2")
	items, _ := rdb.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("  LPUSH 后: %v\n", items)

	rdb.RPush(ctx, key, "item3")
	list, _ := rdb.LRange(ctx, key, 0, -1).Result()
	fmt.Printf("  RPUSH 后: %v\n", list)

	// 2. LPOP & RPOP
	left, _ := rdb.LPop(ctx, key).Result()
	fmt.Printf("  LPOP: %s\n", left)
	right, _ := rdb.RPop(ctx, key).Result()
	fmt.Printf("  RPOP: %s\n", right)

	// 3. LLEN
	length, _ := rdb.LLen(ctx, key).Result()
	fmt.Printf("  LLEN: %d\n", length)

	// 4. BLPOP (阻塞弹出)
	go func() {
		time.Sleep(1 * time.Second)
		rdb.RPush(ctx, "blkey", "task1")
	}()
	fmt.Print("  BLPOP 等待中...")
	result, _ := rdb.BLPop(ctx, 3*time.Second, "blkey").Result()
	fmt.Printf(" 收到: %v\n", result[1])

	// 清理
	rdb.Del(ctx, key, "blkey")
}
