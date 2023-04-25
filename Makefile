run:
	go run main.go

migrate-up:
	migrate -path database/migrations/ -database "postgresql://root:Password123@127.0.0.1:5432/services?sslmode=disable&search_path=public" -verbose up

migrate-down:
	migrate -path database/migrations/ -database "postgresql://root:Password123@127.0.0.1:5432/services?sslmode=disable&search_path=public" -verbose down

mock:
	mockery --dir internal --output internal/mocks --all --keeptree

test:
	ENV=test go test -race -coverprofile coverage.cov -cover ./... && go tool cover -func coverage.cov

test_coverage:
	ENV=test go test ./... -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html
