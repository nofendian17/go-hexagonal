package redis

func (r *Repository) Delete(key string) error {
	if err := r.client.Del(r.ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
