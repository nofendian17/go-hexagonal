.PHONY: migrate-up migrate-down mock test test_coverage

CONFIG_FILE := configs/config.json
CONNECTION_STRING := $(shell jq -r '.database.pgsql | "postgresql://\(.username):\(.password)@\(.host):\(.port)/\(.database)?sslmode=disable&search_path=\(.schema)"' $(CONFIG_FILE))

run:
	go run main.go

migrate-up:
	migrate -path database/migrations/ -database "$(CONNECTION_STRING)" -verbose up

migrate-down:
	migrate -path database/migrations/ -database "$(CONNECTION_STRING)" -verbose down

seed:
	go run database/seeds/seed.go -database "$(CONNECTION_STRING)"

mock:
	mockery --dir internal --output internal/mocks --all --keeptree

test:
	ENV=test go test -race -coverprofile coverage.cov -cover ./... && go tool cover -func coverage.cov

test_coverage:
	ENV=test go test ./... -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html
