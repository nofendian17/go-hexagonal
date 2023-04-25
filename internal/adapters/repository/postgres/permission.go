package postgres

import (
	"errors"
	"user-svc/internal/core/domain"
)

func (r *Repository) CreatePermission(permission *domain.Permission) error {
	query := "INSERT INTO permissions (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(permission.Id, permission.Name, permission.CreatedAt, permission.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePermission(permission *domain.Permission) error {
	query := "UPDATE permissions SET name = $1, updated_at = $2 WHERE id = $3"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(permission.Name, permission.UpdatedAt, permission.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}

func (r *Repository) DeletePermission(id string) error {
	query := "DELETE FROM permissions WHERE id = $1"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}

	return nil
}

func (r *Repository) GetAllPermission() ([]*domain.Permission, error) {
	query := "SELECT id, name, created_at, updated_at FROM permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := make([]*domain.Permission, 0)
	for rows.Next() {
		var permission domain.Permission
		err := rows.Scan(&permission.Id, &permission.Name, &permission.CreatedAt, &permission.UpdatedAt)
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

func (r *Repository) GetPermissionByID(id string) (*domain.Permission, error) {
	query := "SELECT id, name, created_at, updated_at FROM permissions WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var permission domain.Permission
	err := row.Scan(&permission.Id, &permission.Name, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *Repository) PermissionIsExist(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM permissions WHERE name = $1"
	var count int
	row := r.db.QueryRow(query, name)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0 == true, nil
}

func (r *Repository) GetPermissionByName(name string) (*domain.Permission, error) {
	query := "SELECT id, name, created_at, updated_at FROM roles WHERE name = $1"
	row := r.db.QueryRow(query, name)

	var permission domain.Permission
	err := row.Scan(&permission.Id, &permission.Name, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}
