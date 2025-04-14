package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"time"
)

var Ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisClient{Client: client}
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(Ctx, key).Result()
}

func (r *RedisClient) Set(key string, value string, ttlSeconds int) error {
	return r.Client.Set(Ctx, key, value, time.Duration(ttlSeconds)*time.Minute).Err()
}
