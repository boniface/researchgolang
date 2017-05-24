package redisconnection

import "github.com/go-redis/redis"

func GetConnection() ( *redis.Client, error){
	client := redis.NewClient(&redis.Options{
		Addr:     "10.47.2.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client, nil
}