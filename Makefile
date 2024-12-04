generate:
	go generate ./...

run-migrations-up:
	migrate -path=./migrations/ -database postgres://postgres:example@localhost:5432/postgres?sslmode=disable up

run-migrations-down:
	migrate -path=./migrations/ -database postgres://postgres:example@localhost:5432/postgres?sslmode=disable down

create-migration:
	@test -n "$(name)" || (echo "Error: name is not set. Usage: make create-migration name=<migration_name>"; exit 1)
	migrate create -ext sql -dir ./migrations -seq "$(name)"

install-deps:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
