package load

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

// Redis ...
func Redis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:63790",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}

// GetRedis ...
func GetRedis() *redis.Client {
	if rdb == nil {
		panic("redis客户端未初始化")
	}
	return rdb
}
