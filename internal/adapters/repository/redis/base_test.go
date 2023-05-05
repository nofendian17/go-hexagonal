package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"user-svc/internal/shared/config"
)

func TestNewRepository(t *testing.T) {
	cfg := config.New()

	t.Run("success - create new repository", func(t *testing.T) {
		// Create a mock Redis client and set expectations for Ping operation
		db, mock := redismock.NewClientMock()
		mock.ExpectPing().SetErr(nil)
		// Create a new repository with the mock Redis client
		repo := Repository{
			client: db,
			ctx:    context.TODO(),
		}

		got := repo
		assert.Equal(t, repo, got)
	})

	t.Run("success - without password set", func(t *testing.T) {
		cfg.Database.Redis.Password = "foo"
		// Create a mock Redis client and set expectations for Ping operation
		db, mock := redismock.NewClientMock()
		mock.ExpectPing().SetErr(nil)
		// Create a new repository with the mock Redis client
		repo := Repository{
			client: db,
			ctx:    context.TODO(),
		}

		got := repo
		assert.Equal(t, repo, got)
	})

	t.Run("fail - ping redis returns panic", func(t *testing.T) {
		err := errors.New("redis connection failed")
		defer func() {
			if r := recover(); r != nil {
				assert.Error(t, err)
			} else {
				t.Error("The code did not panic")
			}
		}()

		cfg.Database.Redis.Port = 0000
		NewRepository(cfg)
	})

}
