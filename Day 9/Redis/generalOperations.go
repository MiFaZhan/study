package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func generalOperations(rdb *redis.Client) {
	//检查键是否存在
	exists, _ := rdb.Exists(ctx, "key").Result()
	fmt.Println("Key exists:", exists) // 1 存在，0不存在

	// 设置键过期时间为5秒
	rdb.Expire(ctx, "key", 5*time.Second)

	// 获取键的 TTL（剩余生存时间）
	ttl, _ := rdb.TTL(ctx, "key").Result()
	fmt.Println("Key TTL:", ttl) // 输出持续时间，-1 永不过期，-2 键不存在

	// 删除键key1和key2
	rdb.Del(ctx, "key1", "key2")

	// 使用 Pipeline 批量执行命令，减少 RTT (往返时间)
	pipe := rdb.Pipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Pipeline counter:", incr.Val())

	//事务（MULTI/EXEC)
	//Redis 事务不支持回滚，只是保证命令连续执行
	tx := rdb.TxPipeline()
	tx.Set(ctx, "txkey", "txval", 0)
	tx.Incr(ctx, "tx_counter")
	_, err = tx.Exec(ctx)
}
