package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func StringOperations(rdb *redis.Client) {
	// 1.设置值
	err := rdb.Set(ctx, "name", "张三", 0).Err()
	if err != nil {
		println(err)
	}

	//2.获取值
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		println(err)
	}
	println("name:", val)

	//3.设置带过期时间的值
	err = rdb.Set(ctx, "session", "123456", 10*time.Minute).Err()

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
