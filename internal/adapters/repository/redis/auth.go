package redis

import (
	"encoding/json"
	"time"
	"user-svc/internal/core/domain"
)

func (r *Repository) SaveToken(key string, tokenInfo *domain.TokenInfo, expiration time.Duration) error {
	data, err := json.Marshal(tokenInfo)
	if err != nil {
		return err
	}

	err = r.Set(key, data, expiration)
	if err != nil {
		return err
	}
	return nil
}
