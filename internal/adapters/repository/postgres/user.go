package postgres

import (
	"errors"
	"user-svc/internal/core/domain"
)

func (r *Repository) Create(user *domain.User) error {
	query := "INSERT INTO users (id, name, email, salt, password, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.Salt, user.Password, user.Active, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Update(user *domain.User) error {
	query := "UPDATE users SET name = $1, email = $2, salt = $3, password = $4, active = $5, updated_at = $6 WHERE id = $7"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Salt, user.Password, user.Active, user.UpdatedAt, user.Id)
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

func (r *Repository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = $1"
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

func (r *Repository) Users() ([]*domain.User, error) {
	query := "SELECT id, name, email, salt, password, active, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Salt, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) UserByID(id string) (*domain.User, error) {
	query := "SELECT id, name, email, active, created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UserByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name, email, active, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Exist(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	row := r.db.QueryRow(query, email)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0 == true, nil
}
