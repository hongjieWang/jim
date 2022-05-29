package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jim/services/cache"
)

const version = "v1.0"

func main() {
	key := "string:zset"
	set := []redis.Z{
		{Score: 80, Member: "Java"},
		{Score: 90, Member: "Python"},
		{Score: 95, Member: "Golang"},
		{Score: 98, Member: "PHP"},
	}
	err := cache.ZAdd(key, set)
	if err != nil {
		fmt.Println(err)
	}
	scores, _ := cache.ZRevRangeWithScores(key, 0, 2)
	fmt.Println(scores)
	cache.ZIncrBy(key, 5, "Golang")
	scores, _ = cache.ZRevRangeWithScores(key, 0, 2)
	fmt.Println("加分后----")
	fmt.Println(scores)
}
