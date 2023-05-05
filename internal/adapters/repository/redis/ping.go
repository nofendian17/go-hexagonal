package redis

func (r *Repository) Ping() error {
	if err := r.client.Ping(r.ctx).Err(); err != nil {
		return err
	}
	return nil
}
