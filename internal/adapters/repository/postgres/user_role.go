package postgres

import (
	"fmt"
	"strings"
	"user-svc/internal/core/domain"
)

func (r *Repository) GetUserRoles(userID string) ([]*domain.Role, error) {
	query := `
		SELECT r.id, r.name
		FROM user_role ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = $1
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := make([]*domain.Role, 0)
	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *Repository) AddUserRoles(userID string, roles []string) error {
	// Start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	// Build the query string with placeholders for the role IDs
	valueStrings := make([]string, 0, len(roles))
	valueArgs := make([]interface{}, 0, len(roles)*2)
	for i, roleID := range roles {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, userID, roleID)
	}
	query := "INSERT INTO user_role (user_id, role_id) VALUES " + strings.Join(valueStrings, ",")

	// Execute the query with the role IDs as arguments
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RemoveUserRoles(userID string, roles []string) error {
	// Start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	// Build the query string with placeholders for the role IDs
	valueStrings := make([]string, 0, len(roles))
	valueArgs := make([]interface{}, 0, len(roles)*2)
	for i, roleID := range roles {
		valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+2))
		valueArgs = append(valueArgs, roleID)
	}
	query := "DELETE FROM user_role WHERE user_id = $1 AND role_id IN (" + strings.Join(valueStrings, ",") + ")"

	// Execute the query with the role IDs as arguments
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(append([]interface{}{userID}, valueArgs...)...)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
