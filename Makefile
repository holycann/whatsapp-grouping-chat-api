build:
	@go build -o bin/WGCA main.go

test:
	@go test -v ./...

run:build
	@./bin/WGCA

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run db/migrate/migrate.go up

migrate-down:
	@go run db/migrate/migrate.go down

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down