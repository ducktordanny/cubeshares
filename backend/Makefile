build:
	@go build -o bin/cubeshares-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/cubeshares-api

# mac os
install-migrate:
	@brew install golang-migrate

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

db-up:
	@docker compose up -d

db-down:
	@docker compose down -v
	echo "Delete the stopped container if you don't need that anymore."

start: db-up migrate-up run
