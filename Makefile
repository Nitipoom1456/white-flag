GOOSE_DRIVER        ?= postgres
GOOSE_DBSTRING      ?= postgres://whiteflag:whiteflag@localhost:5432/whiteflag?sslmode=disable
GOOSE_MIGRATION_DIR ?= migrations/script

export GOOSE_DRIVER
export GOOSE_DBSTRING
export GOOSE_MIGRATION_DIR

create_migrate_%:
	goose create $* sql

migrate_up:
	goose up

migrate_down:
	goose down

fmt:
	go fmt ./...

run:
	go run cmd/main/main.go
