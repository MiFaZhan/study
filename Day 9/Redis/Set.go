package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func SetOperations(rdb *redis.Client) {
	fmt.Println("\n========== Set 操作 ==========")

	key1 := "set1"
	key2 := "set2"
	rdb.Del(ctx, key1, key2)

	// 1. SADD
	rdb.SAdd(ctx, key1, "a", "b", "c")
	rdb.SAdd(ctx, key2, "b", "c", "d")

	// 2. SMEMBERS
	members, _ := rdb.SMembers(ctx, key1).Result()
	fmt.Printf("  SMEMBERS set1: %v\n", members)

	// 3. SISMEMBER
	isMember, _ := rdb.SIsMember(ctx, key1, "a").Result()
	fmt.Printf("  SISMEMBER set1 'a': %v\n", isMember)

	// 4. SCARD
	count, _ := rdb.SCard(ctx, key1).Result()
	fmt.Printf("  SCARD set1: %d\n", count)

	// 5. 集合运算
	inter, _ := rdb.SInter(ctx, key1, key2).Result()
	fmt.Printf("  SINTER (交集): %v\n", inter)

	union, _ := rdb.SUnion(ctx, key1, key2).Result()
	fmt.Printf("  SUNION (并集): %v\n", union)

	diff, _ := rdb.SDiff(ctx, key1, key2).Result()
	fmt.Printf("  SDIFF (差集 set1-set2): %v\n", diff)

	// 清理
	rdb.Del(ctx, key1, key2)
}
