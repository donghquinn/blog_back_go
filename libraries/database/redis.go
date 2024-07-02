package database

import (
	"context"
	"log"
	"time"

	"github.com/donghquinn/blog_back_go/configs"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisInstance() (*redis.Client, error) {
	 redisConfig := configs.RedisConfig

	redisInstance := redis.NewClient(&redis.Options{
		Addr: redisConfig.Addr,
		// Username: redisConfig.UserName,
		Password: redisConfig.Password,
		DB: 0,
	})

	_, pingError := redisInstance.Ping(ctx).Result()

	if pingError != nil {
		log.Printf("[REDIS] Ping Redis Error: %v", pingError)
		return nil, pingError
	}

	return redisInstance, nil
}

func Set(rdb *redis.Client, key string, token string) error {

	// var item = map[string]string {
	// 	objKey: token}

	// tokenItem := types.RedisToken{
	// 	TokenItem: item,
	// }

	expireDuration := 3 * time.Hour
	setErr := rdb.Set(ctx, key, token, expireDuration).Err()

	if setErr != nil {
		log.Printf("[REDIS] Key Set Error: %v", setErr)

		return setErr
	}

	return nil
}

func Get(rdb *redis.Client, key string) (string, error) {
	getItem, getErr := rdb.Get(ctx, key).Result()

	switch {
		case getErr == redis.Nil:
			log.Printf("[REDIS] No Value Found")
			return "", nil

		case getErr != nil:
			log.Printf("[REDIS] Get Key Error: %v", getErr)
			return "", getErr

		default:
			return getItem, nil
	}
}

func GetAll(rdb *redis.Client, key string) (string, error) {
	getItemList, getErr := rdb.Get(ctx, key).Result()

	if getErr != nil {
		log.Printf("[REDIS] Get Key Error: %v", getErr)
		return "", getErr
	} 

	return getItemList, nil
}

func Delete(rdb *redis.Client, key string, objKey string) error {
	deleteErr := rdb.Del(ctx, key).Err()

	if deleteErr != nil {
		log.Printf("[REDIS] Delete Key Error: %v", deleteErr)
		return deleteErr
	}

	return nil
}

func RedisLoginSet(rdb *redis.Client, sessionId string, email string, userStatus string, userId string, blogId string) error {
	sessionInfo := map[string]string {
		"email": email,
		"userStatus": userStatus,
		"userId": userId,
		"blogId": blogId}

	var ctx = context.Background()

	var data []interface{}

    for key, value := range sessionInfo {
        data = append(data, key, value)
    }

	err := rdb.HSet(ctx, sessionId, data...).Err()
 
    if err != nil {
		log.Printf("[REDIS] Set Value Error: %v", err)
        return err
    }	

	return nil
}

func RedisLoginGet(rdb *redis.Client, sessionId string) (map[string]string, error) {
	var ctx = context.Background()

	getItem, getErr := rdb.HGetAll(ctx, sessionId).Result()
 
	switch {
		case getErr == redis.Nil:
			log.Printf("[REDIS] No Value Found")
			return nil, nil

		case getErr != nil:
			log.Printf("[REDIS] Get Key Error: %v", getErr)
			return nil, getErr

		default:
			return getItem, nil
	}
}