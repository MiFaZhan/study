package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func HashOperations(rdb *redis.Client) {
	var err error

	//1.设置单个字段
	err = rdb.HSet(ctx, "user", "name", "张三").Err()

	//2.设置多个字段
	err = rdb.HSet(ctx, "user:1", map[string]interface{}{
		"name":  "张三",
		"age":   25,
		"email": "zhangsan@example.com",
	}).Err()

	//3.获取单个字段
	name, err := rdb.HGet(ctx, "user:1", "name").Result()
	fmt.Println("Name:", name)

	//4.获取多个字段
	vals, err := rdb.HMGet(ctx, "user:1", "age", "email").Result()
	fmt.Println("Vals:", vals)

	//5.获取所有字段
	all, err := rdb.HGetAll(ctx, "user:1").Result()
	fmt.Println("HGetAll:", all)

	//6.获取所有字段名
	keys, err := rdb.HKeys(ctx, "user:1").Result()
	fmt.Println("HKeys:", keys)

	//7.获取所有字段值
	values, err := rdb.HVals(ctx, "user:1").Result()
	fmt.Println("HVals:", values)

	//8.判断字段是否存在
	exists, err := rdb.HExists(ctx, "user:1", "age").Result()
	fmt.Println("HExists:", exists)

	//9.删除字段
	err = rdb.HDel(ctx, "user:1", "age").Err()
	fmt.Println("HDel:", err)
}
