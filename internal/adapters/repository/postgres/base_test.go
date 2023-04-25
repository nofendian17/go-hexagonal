package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"user-svc/internal/shared/config"
)

func TestNewRepository(t *testing.T) {
	t.Run("Positive Test Case", func(t *testing.T) {
		// CreateUser a mock DB connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error creating mock DB connection: %v", err)
		}
		defer db.Close()

		// Expect a ping to the DB to succeed
		mock.ExpectPing()

		// CreateUser a new config with the DB host, port, etc.
		cfg := config.New()

		// CreateUser a new repository using the mock DB connection
		repo := NewRepository(cfg)
		repo.db = db

		// Ensure that the DB connection was established successfully
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("Error verifying mock expectations: %v", err)
		}
	})

	//t.Run("Negative Test Case", func(t *testing.T) {
	//	// CreateUser a mock DB connection
	//	db, mock, err := sqlmock.New()
	//	if err != nil {
	//		t.Fatalf("Error creating mock DB connection: %v", err)
	//	}
	//	defer db.Close()
	//
	//	// Expect a ping to the DB to return an error
	//	mock.ExpectPing().WillReturnError(errors.New("db ping failure"))
	//
	//	// CreateUser a new config with the DB host, port, etc.
	//	cfg := config.New()
	//
	//	// CreateUser a new repository using the mock DB connection
	//	repo := NewRepository(cfg)
	//	repo.db = db
	//
	//	// Ensure that the error was returned by the Ping() method
	//	if err := mock.ExpectationsWereMet(); err != nil {
	//		t.Fatalf("Error verifying mock expectations: %v", err)
	//	}
	//
	//	// Ensure that the repository was not created due to the ping failure
	//	if repo != nil {
	//		t.Fatalf("Expected repository to be nil due to ping failure, but got %v", repo)
	//	}
	//})
}

func Test_dsn(t *testing.T) {
	cfg := config.New()
	cfg.Database.Pgsql.Host = "127.0.0.1"
	cfg.Database.Pgsql.Port = 5432
	cfg.Database.Pgsql.Username = "foo"
	cfg.Database.Pgsql.Password = "secret"
	cfg.Database.Pgsql.Database = "bar"
	cfg.Database.Pgsql.Schema = "public"

	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{cfg: cfg},
			want: "postgresql://foo:secret@127.0.0.1:5432/bar?sslmode=disable&search_path=public",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dsn(tt.args.cfg); got != tt.want {
				t.Errorf("dsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
