package postgres

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"user-svc/internal/core/domain"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %v", err)
	}
	defer db.Close()

	repo := &Repository{db}

	user := &domain.User{
		Id:        "1",
		Name:      "test",
		Email:     "test@mail.com",
		Salt:      "xxx",
		Password:  "xxx",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `^INSERT INTO (.+) VALUES (.+)`

	t.Run("prepare statement fails", func(t *testing.T) {
		mock.ExpectPrepare(query).
			WillReturnError(fmt.Errorf("failed to prepare statement"))

		err = repo.CreateUser(user)
		assert.Error(t, err)
	})

	t.Run("execution of statement fails", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(user.Id, user.Name, user.Email, user.Salt, user.Password, user.Active, user.CreatedAt, user.UpdatedAt).
			WillReturnError(fmt.Errorf("failed to execute statement"))

		err = repo.CreateUser(user)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(user.Id, user.Name, user.Email, user.Salt, user.Password, user.Active, user.CreatedAt, user.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = repo.CreateUser(user)
		assert.NoError(t, err)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := &Repository{db}

	user := &domain.User{
		Id:        "123",
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Salt:      "abcd1234",
		Password:  "hashed_password",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "UPDATE (.+) SET (.+)"

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).
			WithArgs(user.Name, user.Email, user.Salt, user.Password, user.Active, user.UpdatedAt, user.Id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = r.UpdateUser(user)
		assert.NoError(t, err)
	})

	t.Run("no rows affected", func(t *testing.T) {
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).
			WithArgs(user.Name, user.Email, user.Salt, user.Password, user.Active, user.UpdatedAt, user.Id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err = r.UpdateUser(user)
		assert.Error(t, err)

		assert.True(t, err.Error() == errors.New("no rows were affected").Error())
	})

	t.Run("error preparing statement", func(t *testing.T) {
		expectedErr := errors.New("failed to prepare statement")
		mock.ExpectPrepare(query).WillReturnError(expectedErr)

		err = r.UpdateUser(user)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("error executing statement", func(t *testing.T) {
		expectedErr := errors.New("failed to execute statement")
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).WillReturnError(expectedErr)

		err = r.UpdateUser(user)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("error getting rows affected", func(t *testing.T) {
		expectedErr := errors.New("failed to get rows affected")
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).
			WithArgs(user.Name, user.Email, user.Salt, user.Password, user.Active, user.UpdatedAt, user.Id).
			WillReturnError(expectedErr)

		err = r.UpdateUser(user)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %v", err)
	}
	defer db.Close()

	repo := &Repository{db}
	query := "DELETE FROM (.+) WHERE (.+)"
	// Test case: prepare statement fails
	mock.ExpectPrepare(query).
		WillReturnError(fmt.Errorf("failed to prepare statement"))

	err = repo.DeleteUser("1")
	assert.Error(t, err)

	// Test case: execution of statement fails
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs("1").
		WillReturnError(fmt.Errorf("failed to execute statement"))

	err = repo.DeleteUser("1")
	assert.Error(t, err)

	// Test case: no rows affected
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DeleteUser("1")
	assert.Error(t, err)
	assert.True(t, err.Error() == errors.New("no rows were affected").Error())

	// Test case: success
	mock.ExpectPrepare(query).
		ExpectExec().
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteUser("1")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_Users(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %v", err)
	}
	defer db.Close()

	repo := &Repository{db}

	expectedUsers := []*domain.User{
		{
			Id:        "1",
			Name:      "test1",
			Email:     "test1@mail.com",
			Salt:      "xxx",
			Password:  "xxx",
			Active:    true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        "2",
			Name:      "test2",
			Email:     "test2@mail.com",
			Salt:      "yyy",
			Password:  "yyy",
			Active:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "salt", "password", "active", "created_at", "updated_at"}).
		AddRow(expectedUsers[0].Id, expectedUsers[0].Name, expectedUsers[0].Email, expectedUsers[0].Salt, expectedUsers[0].Password, expectedUsers[0].Active, expectedUsers[0].CreatedAt, expectedUsers[0].UpdatedAt).
		AddRow(expectedUsers[1].Id, expectedUsers[1].Name, expectedUsers[1].Email, expectedUsers[1].Salt, expectedUsers[1].Password, expectedUsers[1].Active, expectedUsers[1].CreatedAt, expectedUsers[1].UpdatedAt)

	// Test case: successfully retrieve users
	mock.ExpectQuery("^SELECT").WillReturnRows(rows)

	users, err := repo.GetAllUsers()
	require.NoError(t, err)
	require.Equal(t, len(expectedUsers), len(users))

	for i, expectedUser := range expectedUsers {
		assert.Equal(t, expectedUser.Id, users[i].Id)
		assert.Equal(t, expectedUser.Name, users[i].Name)
		assert.Equal(t, expectedUser.Email, users[i].Email)
		assert.Equal(t, expectedUser.Salt, users[i].Salt)
		assert.Equal(t, expectedUser.Password, users[i].Password)
		assert.Equal(t, expectedUser.Active, users[i].Active)
		assert.Equal(t, expectedUser.CreatedAt.Unix(), users[i].CreatedAt.Unix())
		assert.Equal(t, expectedUser.UpdatedAt.Unix(), users[i].UpdatedAt.Unix())
	}

	// Test case: failed query
	expectedErr := fmt.Errorf("some error")
	mock.ExpectQuery("^SELECT").WillReturnError(expectedErr)

	users, err = repo.GetAllUsers()
	require.Error(t, err)
	require.Nil(t, users)
	assert.Equal(t, expectedErr, err)

	// Test case: failed row scan
	rows = sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "test")

	mock.ExpectQuery("^SELECT").WillReturnRows(rows)

	users, err = repo.GetAllUsers()
	require.Error(t, err)
	require.Nil(t, users)
	assert.Contains(t, err.Error(), "sql: expected 2 destination arguments in Scan, not 8")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_UserByID(t *testing.T) {
	// Initialize test data and repository
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %v", err)
	}
	defer db.Close()

	repo := &Repository{db}

	id := "123"
	expectedUser := &domain.User{
		Id:        id,
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up mock database response
	query := "SELECT (.+) FROM users WHERE (.+)"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "updated_at"}).
		AddRow(expectedUser.Id, expectedUser.Name, expectedUser.Email, expectedUser.Active, expectedUser.CreatedAt, expectedUser.UpdatedAt)
	mock.ExpectQuery(query).
		WithArgs(id).
		WillReturnRows(rows)

	// Call UserByID method
	user, err := repo.GetUserByID(id)

	// Check for errors and validate user returned by method
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	// Assert that all mock expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserByEmail(t *testing.T) {
	// CreateUser a mock DB connection for the test
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB connection: %s", err)
	}
	defer db.Close()

	// CreateUser the repository with the mock DB connection
	repo := &Repository{db: db}

	// Define the expected user and row data
	expectedUser := &domain.User{
		Id:        "1",
		Name:      "Test GetUser",
		Email:     "test@example.com",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	row := sqlmock.NewRows([]string{"id", "name", "email", "active", "salt", "password", "created_at", "updated_at"}).
		AddRow(expectedUser.Id, expectedUser.Name, expectedUser.Email, expectedUser.Active, expectedUser.Salt, expectedUser.Password, expectedUser.CreatedAt, expectedUser.UpdatedAt)

	// Set up the mock DB to return the expected row data
	mock.ExpectQuery("^SELECT (.+) FROM users WHERE email = (.+)$").
		WithArgs(expectedUser.Email).
		WillReturnRows(row)

	// Call the method being tested
	actualUser, err := repo.GetUserByEmail(expectedUser.Email)
	if err != nil {
		t.Fatalf("Unexpected error from UserByEmail: %s", err)
	}

	// Verify that the expected user was returned
	assert.Equal(t, expectedUser, actualUser)

	// Verify that all expected mock DB calls were made
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations: %s", err)
	}
}

func TestExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	repo := &Repository{db: db}

	email := "test@example.com"

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT COUNT(.+) FROM users").WithArgs(email).WillReturnRows(rows)

	exist, err := repo.UserIsExist(email)
	assert.NoError(t, err)
	assert.True(t, exist)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
