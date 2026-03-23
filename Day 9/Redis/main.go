package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// 创建一个空的上下文（context）变量，用于在 Go 程序中传递请求范围的数据、取消信号和超时等。
var ctx = context.Background()

func main() {
	// 创建 redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10, // 连接池大小
		MinIdleConns: 5,  // 最小空闲连接数
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("========================================")
	fmt.Printf("✓ Redis 连接成功 (PING: %s)\n", pong)
	fmt.Println("========================================")

	StringOperations(rdb)
	SetOperations(rdb)
	ListOperations(rdb)
	HashOperations(rdb)
	generalOperations(rdb)

	fmt.Println("\n========================================")
	fmt.Println("✓ 所有操作完成")
	fmt.Println("========================================")
}
