package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"

	"user-svc/internal/shared/config"
)

type Repository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRepository(cfg *config.Config) *Repository {
	ctx := context.Background()
	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Database.Redis.Host, cfg.Database.Redis.Port),
		DB:   cfg.Database.Redis.DB,
	}
	if cfg.Database.Redis.Password != "" {
		options.Password = cfg.Database.Redis.Password
	}

	client := redis.NewClient(options)
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(fmt.Errorf("redis ping failure: %v", err))
	}

	return &Repository{
		client: client,
		ctx:    ctx,
	}
}

func (r *Repository) MarshalBinary(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (r *Repository) UnmarshalBinary(data []byte) (result interface{}, err error) {
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}
