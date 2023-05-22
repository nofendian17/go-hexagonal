package data

import (
	"database/sql"
	"fmt"
	"user-svc/internal/shared/config"
)

type SeederInterface interface {
	Run() (err error)
	Clear(table string) (err error)
}

type Seeder struct {
	db     *sql.DB
	config *config.Config
}

func NewSeeder(db *sql.DB, config *config.Config) *Seeder {
	return &Seeder{
		db:     db,
		config: config,
	}
}

func (s *Seeder) Run() (err error) {
	if err := s.userSeeder(); err != nil {
		return err
	}
	if err := s.roleSeeder(); err != nil {
		return err
	}
	if err := s.userRoleSeeder(); err != nil {
		return err
	}
	if err := s.permissionSeeder(); err != nil {
		return err
	}
	if err := s.rolePermissionSeeder(); err != nil {
		return err
	}
	return nil
}

func (s *Seeder) Clear(table string) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s`, table)
	_, err = s.db.Query(query)
	if err != nil {
		return err
	}
	return err
}
