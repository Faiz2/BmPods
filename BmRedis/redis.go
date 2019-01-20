package BmRedis

import (
	"errors"
	"fmt"
	"time"

	bmconfig "github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {

	redisConfigPath := "Resources/redisconfig.json"
	redisConfig := bmconfig.BMGetConfigMap(redisConfigPath)

	host := redisConfig["Host"].(string)
	port := redisConfig["Port"].(string)
	addr := host + ":" + port
	password := redisConfig["Password"].(string)
	db := int(redisConfig["DB"].(float64))

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	return client
}

func CheckToken(token string) error {
	client := GetRedisClient()
	defer client.Close()

	_, err := client.Get(token).Result()
	if err == redis.Nil {
		return errors.New("token not exist")
	} else if err != nil {
		//panic(err)
		fmt.Println(err.Error())
		return err
	} else {
		return nil
	}
}

func PushToken(token string) error {
	client := GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Incr(token)
	pipe.Expire(token, 365*24*time.Hour)

	_, err := pipe.Exec()
	fmt.Println(token)
	return err
}
