package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"reminder/config"
)

func NewRedis(conf *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	})
}
