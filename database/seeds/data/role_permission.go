package data

import (
	"fmt"
)

func (s *Seeder) rolePermissionSeeder() error {
	table := "role_permission"
	roleTable := "roles"
	permissionTable := "permissions"
	err := s.Clear(table)
	if err != nil {
		return err
	}

	// Define the roles and their corresponding permissions
	rolePermissions := map[string][]string{
		"Admin": {
			"List-User", "View-User", "Create-User", "Update-User", "Delete-User",
			"List-Role", "View-Role", "Create-Role", "Update-Role", "Delete-Role",
			"List-Permission", "View-Permission", "Create-Permission", "Update-Permission", "Delete-Permission",
		},
		"Manager": {
			"List-User", "View-User", "Create-User", "Update-User",
			"List-Role", "View-Role", "Create-Role", "Update-Role",
			"List-Permission", "View-Permission",
		},
		"User": {
			"List-User", "View-User", "Create-User", "Update-User",
		},
		"Guest": {
			"View-User", "View-Role", "View-Permission",
		},
	}

	// Prepare the SQL statement for inserting role permission records
	stmt, err := s.db.Prepare(fmt.Sprintf("INSERT INTO %s (role_id, permission_id) VALUES ($1, $2)", table))
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Query the roles and permissions tables to get the role_id and permission_id values
	roleQuery := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", roleTable)
	permissionQuery := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", permissionTable)

	for role, permissions := range rolePermissions {
		// Get the role_id for the current role
		var roleID string
		err := s.db.QueryRow(roleQuery, role).Scan(&roleID)
		if err != nil {
			return err
		}

		// Insert the role permission records
		for _, permission := range permissions {
			// Get the permission_id for the current permission
			var permissionID string
			err := s.db.QueryRow(permissionQuery, permission).Scan(&permissionID)
			if err != nil {
				return err
			}

			_, err = stmt.Exec(roleID, permissionID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
