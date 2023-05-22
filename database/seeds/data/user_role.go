package data

import (
	"fmt"
	"github.com/jaswdr/faker"
)

func (s *Seeder) userRoleSeeder() error {
	userTable := "users"
	roleTable := "roles"
	userRoleTable := "user_role"

	// Clear the user_role table
	err := s.Clear(userRoleTable)
	if err != nil {
		return err
	}

	// Prepare the SQL statement for inserting user role references
	stmt, err := s.db.Prepare(fmt.Sprintf(`INSERT INTO %s (user_id, role_id) VALUES ($1, $2)`, userRoleTable))
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Select data from users and roles tables
	userQuery := fmt.Sprintf(`SELECT id FROM %s`, userTable)
	roleQuery := fmt.Sprintf(`SELECT id FROM %s`, roleTable)

	userRows, err := s.db.Query(userQuery)
	if err != nil {
		return err
	}
	defer userRows.Close()

	roleRows, err := s.db.Query(roleQuery)
	if err != nil {
		return err
	}
	defer roleRows.Close()

	// Store user and role IDs in slices
	var userIDs []string
	var roleIDs []string

	for userRows.Next() {
		var userID string
		err := userRows.Scan(&userID)
		if err != nil {
			return err
		}
		userIDs = append(userIDs, userID)
	}

	for roleRows.Next() {
		var roleID string
		err := roleRows.Scan(&roleID)
		if err != nil {
			return err
		}
		roleIDs = append(roleIDs, roleID)
	}

	// Perform the user role seeding
	fake := faker.New()
	for _, userID := range userIDs {
		// Randomly select a role ID
		roleID := fake.RandomStringElement(roleIDs)

		_, err := stmt.Exec(userID, roleID)
		if err != nil {
			return err
		}
	}

	return nil
}
