package postgres

import (
	"fmt"
	"strings"
	"user-svc/internal/core/domain"
)

func (r *Repository) GetRolePermissions(roleId string) ([]*domain.Permission, error) {
	query := `
		SELECT p.id, p.name
		FROM role_permission rp
		INNER JOIN permissions p ON p.id = rp.permission_id
		WHERE rp.role_id = $1
	`
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := make([]*domain.Permission, 0)
	for rows.Next() {
		var permission domain.Permission
		err := rows.Scan(&permission.Id, &permission.Name)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, &permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *Repository) AddRolePermissions(roleId string, permissions []string) error {
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

	// Build the query string with placeholders for the permission IDs
	valueStrings := make([]string, 0, len(permissions))
	valueArgs := make([]interface{}, 0, len(permissions)*2)
	for i, permissionID := range permissions {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, roleId, permissionID)
	}
	query := "INSERT INTO role_permission (role_id, permission_id) VALUES " + strings.Join(valueStrings, ",")

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

func (r *Repository) RemoveRolePermissions(roleId string, permissions []string) error {
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

	// Build the query string with placeholders for the permission IDs
	valueStrings := make([]string, 0, len(permissions))
	valueArgs := make([]interface{}, 0, len(permissions)*2)
	for i, permissionID := range permissions {
		valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+2))
		valueArgs = append(valueArgs, permissionID)
	}
	query := "DELETE FROM role_permission WHERE role_id = $1 AND permission_id IN (" + strings.Join(valueStrings, ",") + ")"

	// Execute the query with the role IDs as arguments
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(append([]interface{}{roleId}, valueArgs...)...)
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
