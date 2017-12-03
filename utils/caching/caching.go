package caching

import (
	"github.com/go-redis/redis"
	"time"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, expiration time.Duration) error
}

type Redis struct {
	Client *redis.Client
}

func Connect(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

func (r *Redis) Set(key, value string, expiration time.Duration) error {
	return r.Client.Set(key, value, expiration).Err()
}
