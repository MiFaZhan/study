package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func SortedSetOperations(rdb *redis.Client) {
	key := "leaderboard"
	rdb.Del(ctx, key)

	// ZADD
	members := []redis.Z{
		{Score: 100, Member: "player:1"},
		{Score: 250, Member: "player:2"},
		{Score: 150, Member: "player:3"},
	}
	rdb.ZAdd(ctx, key, members...)

	// ZRANGE (从小到大)
	all, _ := rdb.ZRange(ctx, key, 0, -1).Result()
	fmt.Println("All players (asc):", all) // 输出: All players: [player:1 player:3 player:2]

	// ZRANGEWITHSCORES
	allWithScores, _ := rdb.ZRangeWithScores(ctx, key, 0, -1).Result()
	for _, z := range allWithScores {
		fmt.Printf("Player: %s, Score: %.0f\n", z.Member, z.Score)
	}

	// ZREVRANGE (从大到小)
	top2, _ := rdb.ZRevRange(ctx, key, 0, 1).Result()
	fmt.Println("Top 2 players:", top2) // 输出: Top 2 players: [player:2 player:3]

	// ZSCORE
	score, _ := rdb.ZScore(ctx, key, "player:1").Result()
	fmt.Println("Player:1 score:", score) // 输出: Player:1 score: 100

	// ZRANK & ZREVRANK
	rank, _ := rdb.ZRank(ctx, key, "player:2").Result()               // 从小到大排名，从0开始
	revRank, _ := rdb.ZRevRank(ctx, key, "player:2").Result()         // 从大到小排名
	fmt.Printf("Player:2 rank: %d (asc), %d (desc)\n", rank, revRank) // 输出: Player:2 rank: 2 (asc), 0 (desc)

	// ZINCRBY
	newScore, _ := rdb.ZIncrBy(ctx, key, 50, "player:1").Result()
	fmt.Println("Player:1 new score:", newScore) // 输出: Player:1 new score: 150

	// ZCOUNT
	count, _ := rdb.ZCount(ctx, key, "100", "200").Result()
	fmt.Println("Players with score between 100 and 200:", count) // 输出:: 2
}
