package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"app/pkg/conf"
)

var (
	// RC redis cache client
	RC *redis.Client
)

func Init(cfg *conf.RedisConfig) (err error) {
	if cfg == nil {
		return
	}

	ctx := context.Background()
	readTimeout := time.Duration(30) * time.Second
	writeTimeout := time.Duration(30) * time.Second
	RC = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		DB:           cfg.Db,
		Password:     cfg.Password,
		PoolSize:     cfg.PoolSize,
		MaxIdleConns: cfg.MaxIdle,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})

	if _, err = RC.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}
