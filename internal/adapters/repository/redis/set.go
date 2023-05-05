package redis

import "time"

func (r *Repository) Set(key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(r.ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}
