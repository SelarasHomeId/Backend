package database

import (
	"fmt"
	"selarashomeid/internal/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func InitRedis() *redis.Client {
	var redisHost = fmt.Sprintf("%s:%s", config.Get().Redis.RedisAddress, config.Get().Redis.RedisPort)
	var redisPassword = config.Get().Redis.RedisPassword

	rdb := newRedisClient(redisHost, redisPassword)
	logrus.Info(fmt.Sprintf("Successfully connected to %s", "REDIS"))

	return rdb
}
