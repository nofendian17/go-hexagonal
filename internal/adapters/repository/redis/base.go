package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"

	"user-svc/internal/shared/config"
)

type Repository struct {
	client     *redis.Client
	ctx        context.Context
	defaultTTL time.Duration
}

func NewRepository(cfg *config.Config) *Repository {
	ctx := context.Background()
	opt := options(cfg)

	client := redis.NewClient(opt)
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Errorf("redis ping failure: %v", err))
	}

	return &Repository{
		client:     client,
		ctx:        ctx,
		defaultTTL: time.Duration(cfg.Database.Redis.Lifetime),
	}
}

func options(cfg *config.Config) *redis.Options {
	addr := fmt.Sprintf("%s:%d", cfg.Database.Redis.Host, cfg.Database.Redis.Port)
	if cfg.App.Debug {
		fmt.Println(fmt.Sprintf("Trying connect redis with %s on DB %d", addr, cfg.Database.Redis.DB))
	}

	return &redis.Options{
		Network:               "",
		Addr:                  addr,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Username:              "",
		Password:              cfg.Database.Redis.Password,
		CredentialsProvider:   nil,
		DB:                    cfg.Database.Redis.DB,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	}
}
