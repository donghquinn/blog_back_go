package database

import (
	"context"
	"log"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/donghquinn/blog_back_go/types"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisInstance() *redis.Client {
	 redisConfig := configs.RedisConfig

	redis := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		Password: redisConfig.Password,
	})

	return redis
}

func Set(redis *redis.Client, key string, objKey string, token string) error {

	var item = map[string]string {
		objKey: token}

	tokenItem := types.RedisToken{
		TokenItem: item,
	}

	expireDuration := 3 * time.Hour
	setErr := redis.Set(ctx, key, tokenItem, expireDuration).Err()

	if setErr != nil {
		log.Printf("[REDIS] Key Set Error: %v", setErr)

		return setErr
	}

	return nil
}

func Get(redis *redis.Client, key string, objKey string) (string, error) {
	getItem, getErr := redis.Get(ctx, key).Result()

	if getErr != nil {
		log.Printf("[REDIS] Get Key Error: %v", getErr)
		return "", getErr
	} 

	return getItem, nil
}

func GetAll(redis *redis.Client, key string, objKey string) (string, error) {
	getItem, getErr := redis.Get(ctx, key).Result()

	if getErr != nil {
		log.Printf("[REDIS] Get Key Error: %v", getErr)
		return "", getErr
	} 

	return getItem, nil
}

func Delete(redis *redis.Client, key string, objKey string) error {
	deleteErr := redis.Del(ctx, key).Err()

	if deleteErr != nil {
		log.Printf("[REDIS] Delete Key Error: %v", deleteErr)
		return deleteErr
	}

	return nil
}