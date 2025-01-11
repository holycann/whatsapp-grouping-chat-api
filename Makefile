build:
	@go build -o bin/myapi cmd/main.go

test:
	@go test -v ./...

run:build
	@./bin/myapi

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrate.go up

migrate-down:
	@go run cmd/migrate/migrate.go down