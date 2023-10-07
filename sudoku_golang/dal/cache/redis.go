package cache

import (
	"errors"

	redis "github.com/redis/go-redis/v9"
)

var (
	RedisDB *redis.Client
)

func Init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "43.136.122.18:6379",
		Password: "147258",
		DB:       3,
	})
	if RedisDB == nil {
		panic(errors.New("[redis init error]"))
	}
}
