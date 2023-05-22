package data

import (
	"fmt"
	"github.com/jaswdr/faker"
	"time"
)

func (s *Seeder) roleSeeder() (err error) {
	table := "roles"
	err = s.Clear(table)

	if err != nil {
		return err
	}

	fake := faker.New()
	now := time.Now()
	roles := []string{"Admin", "Manager", "User", "Guest"}

	for i := 0; i < len(roles); i++ {
		query := fmt.Sprintf("INSERT INTO %s (id, name, active, created_at, updated_at) values ($1, $2, $3, $4, $5)", table)
		stmt, err := s.db.Prepare(query)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(fake.UUID().V4(), roles[i], true, now, now)
		if err != nil {
			return err
		}
	}
	return err
}
