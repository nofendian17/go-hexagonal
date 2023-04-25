package postgres

import (
	"errors"
	"user-svc/internal/core/domain"
)

func (r *Repository) CreateRole(role *domain.Role) error {
	query := "INSERT INTO roles (id, name, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(role.Id, role.Name, role.Active, role.CreatedAt, role.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateRole(role *domain.Role) error {
	query := "UPDATE roles SET name = $1, active = $2, updated_at = $3 WHERE id = $4"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(role.Name, role.Active, role.UpdatedAt, role.Id)
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

func (r *Repository) DeleteRole(id string) error {
	query := "DELETE FROM roles WHERE id = $1"
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

func (r *Repository) GetAllRole() ([]*domain.Role, error) {
	query := "SELECT id, name, active, created_at, updated_at FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := make([]*domain.Role, 0)
	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.Id, &role.Name, &role.Active, &role.CreatedAt, &role.UpdatedAt)
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

func (r *Repository) GetRoleByID(id string) (*domain.Role, error) {
	query := "SELECT id, name, active, created_at, updated_at FROM roles WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var role domain.Role
	err := row.Scan(&role.Id, &role.Name, &role.Active, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *Repository) RoleIsExist(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM roles WHERE name = $1"
	var count int
	row := r.db.QueryRow(query, name)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0 == true, nil
}

func (r *Repository) GetRoleByName(name string) (*domain.Role, error) {
	query := "SELECT id, name, active, created_at, updated_at FROM roles WHERE name = $1"
	row := r.db.QueryRow(query, name)

	var role domain.Role
	err := row.Scan(&role.Id, &role.Name, &role.Active, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
