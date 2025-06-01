build:
	@go build -o /bin/cubeit-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/cubeit-api

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
	@docker compose up

db-down:
	@docker compose down
	echo "Delete the stopped container if you don't need that anymore."
