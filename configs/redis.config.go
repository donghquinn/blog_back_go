package configs

import "os"

type RedisConfigStruct struct {
	Addr string
	Password string
	// DB int
}

var RedisConfig RedisConfigStruct

func SetRedisConfig() {
	RedisConfig.Addr = os.Getenv("REDIS_ADDR")
	RedisConfig.Password = os.Getenv("REDIS_PASSWORD")
	// redisConfig.DB
}