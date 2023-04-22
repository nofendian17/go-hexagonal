run:
	go run main.go

migrate-up:
	migrate -path database/migrations/ -database "postgresql://root:Password123@127.0.0.1:5432/services?sslmode=disable&search_path=public" -verbose up

migrate-down:
	migrate -path database/migrations/ -database "postgresql://root:Password123@127.0.0.1:5432/services?sslmode=disable&search_path=public" -verbose down

mock:
	mockery --dir internal --output internal/mocks --all --keeptree

