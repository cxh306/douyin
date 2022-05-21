package redis

import (
	"github.com/go-redis/redis"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	expireTime = 600
	rdb        *redis.Client
)

func Init() error {
	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: ""})
	_, err := rdb.Ping().Result()
	return err
}
