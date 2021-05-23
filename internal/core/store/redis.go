package store

import (
	"fmt"
	"go-frame/internal/core/setting"
	"time"

	redis "github.com/go-redis/redis/v8"
)

func NewRedisClientPool(setting *setting.RedisSettings) *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d", setting.Host, setting.Port),
			//MaxRetries:         0,
			//MinRetryBackoff:    0,
			//MaxRetryBackoff:    0,
			DialTimeout:        time.Millisecond * 500,
			ReadTimeout:        time.Millisecond * 500,
			WriteTimeout:       time.Millisecond * 500,
			PoolSize:           5,
			MinIdleConns:       2,
			MaxConnAge:         time.Second * 20,
			PoolTimeout:        time.Millisecond * 500,
			IdleTimeout:        time.Second * 5,
			IdleCheckFrequency: time.Second * 1,
		},
	)
	return client
}
