package data

import (
	"AI_Server/init/conf"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"time"
)

var RDB *redis.Client
var RedisCount int8

func InitRedis() *redis.Client {
	redisConfig := conf.GlobalConfig.Data.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Username:     redisConfig.Username, // default user
		Password:     redisConfig.Password,
		WriteTimeout: redisConfig.WriteTimeout,
		ReadTimeout:  redisConfig.ReadTimeout,
		DB:           0,               // use default DB
		DialTimeout:  1 * time.Second, // 1 second
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Error().Msg("Redis 连接失败")
		panic(err)
	} else {
		log.Info().Msg("Redis 连接成功")
	}
	RDB = rdb
	return rdb
}
