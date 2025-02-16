include .env

ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: install-goose
install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: db-up
db-up:
	docker-compose up -d
	make install-goose
	make migration-up

.PHONY: run
run:
	docker-compose up

.PHONY: install-lint
install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5

lint: install-lint
	golangci-lint run -c .golangci.yaml

test:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./...

go-cover:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	cat profile.cov | grep -v "mock" > prof.cov
	go tool cover -func prof.cov