package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func HashOperations(rdb *redis.Client) {
	fmt.Println("\n========== Hash 操作 ==========")

	// 1. 设置多个字段
	rdb.HSet(ctx, "user:1", map[string]interface{}{
		"name":  "张三",
		"age":   25,
		"email": "zhangsan@example.com",
	})

	// 2. 获取单个字段
	name, _ := rdb.HGet(ctx, "user:1", "name").Result()
	fmt.Printf("  HGET user:1 name: %s\n", name)

	// 3. 获取多个字段
	vals, _ := rdb.HMGet(ctx, "user:1", "age", "email").Result()
	fmt.Printf("  HMGET user:1 age email: %v\n", vals)

	// 4. 获取所有字段
	all, _ := rdb.HGetAll(ctx, "user:1").Result()
	fmt.Printf("  HGETALL user:1: %v\n", all)

	// 5. 判断字段是否存在
	exists, _ := rdb.HExists(ctx, "user:1", "age").Result()
	fmt.Printf("  HEXISTS user:1 age: %v\n", exists)

	// 6. 删除字段
	rdb.HDel(ctx, "user:1", "age")
	fmt.Println("  HDEL user:1 age: 已删除")

	// 清理
	rdb.Del(ctx, "user", "user:1")
}
