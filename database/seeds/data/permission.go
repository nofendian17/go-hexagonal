package data

import (
	"fmt"
	"github.com/jaswdr/faker"
	"time"
)

func (s *Seeder) permissionSeeder() (err error) {
	table := "permissions"
	err = s.Clear(table)

	if err != nil {
		return err
	}

	fake := faker.New()
	now := time.Now()
	permissions := []string{
		"List-User", "View-User", "Create-User", "Update-User", "Delete-User",
		"List-Role", "View-Role", "Create-Role", "Update-Role", "Delete-Role",
		"List-Permission", "View-Permission", "Create-Permission", "Update-Permission", "Delete-Permission",
	}

	for i := 0; i < len(permissions); i++ {
		query := fmt.Sprintf("INSERT INTO %s (id, name, created_at, updated_at) values ($1, $2, $3, $4)", table)
		stmt, err := s.db.Prepare(query)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(fake.UUID().V4(), permissions[i], now, now)
		if err != nil {
			return err
		}
	}
	return err
}
