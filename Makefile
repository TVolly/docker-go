#!make
-include .env

DB_DRIVER ?= postgres
DB_USERNAME ?= postgres
DB_PASSWORD ?= secret
DB_HOST ?= db
DB_PORT ?= 5432
DB_DATABASE ?= $(DB_USERNAME)

MIGRATE_DB_URL = $(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable


help:		## Help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'



up:			## Up application
	@docker-compose up -d

stop:		## Stop application
	@docker-compose stop

down:		## Down application
	@docker-compose down


db:		## Enter to DB
	@docker-compose \
		exec db \
		psql -U $(DB_USERNAME) $(DB_DATABASE)


tests:	## Run tests
	@docker-compose \
		run --rm app \
		go test -v -vet=off ./... 

run:	## Run app
	@docker-compose \
		run --rm app \
		go run ./cmd/main.go 

build:	## Build app
	@docker-compose \
		run --rm app \
		go build -v ./cmd/main.go
	

migrate-create:	## Create migration [-n=name]
	@docker-compose \
		run --rm migrate \
		create -dir ./ -ext sql -seq $(n)

migrate-version: ## Get migration version
	@docker-compose \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ version 

migrate-up:		## Run migrations
	@docker-compose \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ up 

migrate-down:	## Rollback migrations
	@docker-compose \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ down 

migrate-force:	## Force migration [-v=version]
	@docker-compose \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ force $(v)


.DEFAULT_GOAL := help