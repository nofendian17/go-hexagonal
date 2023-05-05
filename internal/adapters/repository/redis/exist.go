package redis

func (r *Repository) Exists(key string) (bool, error) {
	val, err := r.client.Exists(r.ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}
