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

func RedisInstance() (*redis.Client, error) {
	 redisConfig := configs.RedisConfig

	redisInstance := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		Password: redisConfig.Password,
	})

	pingResult, pingError := redisInstance.Ping(ctx).Result()

	if pingError != nil {
		log.Printf("[REDIS] Ping Redis Error: %v", pingError)
		return nil, pingError
	}

	log.Printf("[REDIS] Successfully Connect to Redis: %s", pingResult)

	return redisInstance, nil
}

func Set(redis *redis.Client, key string, token string) error {

	// var item = map[string]string {
	// 	objKey: token}

	// tokenItem := types.RedisToken{
	// 	TokenItem: item,
	// }

	expireDuration := 3 * time.Hour
	setErr := redis.Set(ctx, key, token, expireDuration).Err()

	if setErr != nil {
		log.Printf("[REDIS] Key Set Error: %v", setErr)

		return setErr
	}

	return nil
}

func Get(redis *redis.Client, key string) (string, error) {
	getItem, getErr := redis.Get(ctx, key).Result()

	if getErr != nil {
		log.Printf("[REDIS] Get Key Error: %v", getErr)
		return "", getErr
	}

	return getItem, nil
}

func GetAll(redis *redis.Client, key string) (string, error) {
	getItemList, getErr := redis.Get(ctx, key).Result()

	if getErr != nil {
		log.Printf("[REDIS] Get Key Error: %v", getErr)
		return "", getErr
	} 

	return getItemList, nil
}

func Delete(redis *redis.Client, key string, objKey string) error {
	deleteErr := redis.Del(ctx, key).Err()

	if deleteErr != nil {
		log.Printf("[REDIS] Delete Key Error: %v", deleteErr)
		return deleteErr
	}

	return nil
}

func RedisLoginSet(redis *redis.Client, sessionId string, email string, userStatus string, userId string) error {
	sessionInfo := types.LoginRedisStruct {
		Email: email,
		UserStatus: userStatus,
		UserId: userId}

	var ctx = context.Background()

	err := redis.Set(ctx, sessionId, sessionInfo, time.Hour * 3).Err()
 
    if err != nil {
		log.Printf("[REDIS] Set Value Error: %v", err)
        return err
    }	

	return nil
}