package redis

func (r *Repository) Close() error {
	return r.client.Close()
}
