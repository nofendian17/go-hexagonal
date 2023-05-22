package data

import (
	"encoding/base64"
	"fmt"
	"github.com/jaswdr/faker"
	"time"
	"user-svc/internal/shared/hash"
)

func (s *Seeder) userSeeder() (err error) {
	table := "users"
	password := "123456"
	length := 100

	err = s.Clear(table)
	if err != nil {
		return err
	}

	fake := faker.New()
	hashes := hash.NewHasher(s.config)

	for i := 0; i < length; i++ {
		query := fmt.Sprintf("INSERT INTO %s (id, name, email, active, salt, password, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)", table)
		stmt, err := s.db.Prepare(query)
		if err != nil {
			return err
		}

		salt, err := hashes.GenerateRandomSalt()
		if err != nil {
			return err
		}
		hashedPassword := hashes.HashPassword(password, salt)
		now := time.Now()

		_, err = stmt.Exec(fake.UUID().V4(), fake.Person().Name(), fake.Internet().Email(), true, base64.URLEncoding.EncodeToString(salt), hashedPassword, now, now)
		if err != nil {
			return err
		}
	}
	return err
}
