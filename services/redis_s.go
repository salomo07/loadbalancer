package services

import (
	"context"
	"loadbalancer/config"
	"time"

	"github.com/redis/go-redis/v9"
)

func SaveValueRedis(key string, value string) {
	startTime := time.Now()
	var ctx = context.Background()
	if config.REDIS_CRED == "" {
		config.GetCredRedis()
	}

	opt, _ := redis.ParseURL(config.REDIS_CRED)
	client := redis.NewClient(opt)

	client.Set(ctx, key, value, 0)
	endTime := time.Now()
	responseTime := endTime.Sub(startTime).Nanoseconds()
	print(responseTime)
	print(key + " is saved")
}

func GetValueRedis(key string) string {
	var ctx = context.Background()
	if config.REDIS_CRED == "" {
		config.GetCredRedis()
	}

	opt, _ := redis.ParseURL(config.REDIS_CRED)
	client := redis.NewClient(opt)
	// startTime := time.Now()
	var res = client.Get(ctx, key).Val()
	// endTime := time.Now()
	// responseTime := endTime.Sub(startTime).Nanoseconds()
	// println("Time process : " + string(responseTime))
	return res
}
