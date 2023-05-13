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

func (r *Repository) GetToken(key string) (*domain.TokenInfo, error) {
	tokenInfo := &domain.TokenInfo{}
	result, err := r.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func (r *Repository) DeleteToken(key string) error {
	return r.Delete(key)
}
