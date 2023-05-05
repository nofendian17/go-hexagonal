package ports

import (
	"time"
)

type CacheRepository interface {
	Close() error
	Ping() error
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	MarshalBinary(value interface{}) ([]byte, error)
	UnmarshalBinary(data []byte) (interface{}, error)
}
