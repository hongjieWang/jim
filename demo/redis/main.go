package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jim/services/cache"
	"time"
)

var RedisCache = &redis.Client{}

func init() {
	RedisCache = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "125846whj",
		DB:       0,
	})
	//ping
	pong, err := RedisCache.Ping().Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return
	}
	fmt.Println("ping result:", pong)
}

// SetString SET
func SetString(key, value string) {
	err := RedisCache.Set(key, value, time.Hour).Err()
	if err != nil {
		fmt.Println(err)
	}
}
func Get(key string) string {
	result, err := RedisCache.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func SetUser() {
	key := "string:user"
	user := User{Name: "张三", Age: 28}
	jsonValue, _ := json.Marshal(user)
	if RedisCache.Set(key, jsonValue, time.Hour).Err() != nil {
		fmt.Println("Set Cache Err")
	}

	result, _ := RedisCache.Get(key).Result()
	fmt.Println(result)
	var user2 User
	json.Unmarshal([]byte(result), &user2)
	fmt.Println(user2)
}

func List() {
	key := "string:list"
	err := RedisCache.LPush(key, "A", "B", "C", 20, "D", "E", "F").Err()
	if err != nil {
		fmt.Println("缓存设置错误", err)
	}
	lLen, _ := RedisCache.LLen(key).Result()
	fmt.Printf("集合数据长度：%d\n", lLen)
	lRange, _ := RedisCache.LRange(key, 0, 3).Result()
	fmt.Println(lRange)
}

func Hash() {
	key := "string:hash"
	RedisCache.HSet(key, "name", "张三")
	RedisCache.HSet(key, "phone", "18234554345")
	RedisCache.HSet(key, "age", "28")
	//获取全部hash对象
	all, _ := RedisCache.HGetAll(key).Result()
	fmt.Println(all)
	//修改已存在的字段->如果field已存在，则对值进行覆盖操作。
	RedisCache.HSet(key, "name", "李四")
	RedisCache.HSet(key, "email", "july@163.com")
	all, _ = RedisCache.HGetAll(key).Result()
	fmt.Println(all)

	////获取指定字段
	name, _ := RedisCache.HGet(key, "name").Result()
	fmt.Println(name)
	existsName, _ := RedisCache.HExists(key, "name").Result()
	existsId, _ := RedisCache.HExists(key, "id").Result()
	fmt.Printf("name 字段是否存在 %v\n", existsName)
	fmt.Printf("id 字段是否存在 %v\n", existsId)
	RedisCache.HDel(key, "name")
	existsName, _ = RedisCache.HExists(key, "name").Result()
	fmt.Printf("name 字段是否存在 %v\n", existsName)
	getAll, _ := RedisCache.HGetAll(key).Result()
	fmt.Println(getAll)
}

func ZSet() {
	key := "string:zset"
	set := []redis.Z{
		{Score: 80, Member: "Java"},
		{Score: 90, Member: "Python"},
		{Score: 95, Member: "Golang"},
		{Score: 98, Member: "PHP"},
	}
	err := RedisCache.ZAdd(key, set...)
	if err != nil {
		fmt.Println(err)
	}
	scores, _ := RedisCache.ZRevRangeWithScores(key, 0, 2).Result()
	fmt.Println(scores)
	cache.ZIncrBy(key, 5, "Golang")
	scores, _ = RedisCache.ZRevRangeWithScores(key, 0, 2).Result()
	fmt.Println("加分后----")
	fmt.Println(scores)
}

func main() {
	fmt.Println("Golang 操作Redis Demo")
	ZSet()
}
