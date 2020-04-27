#!make
-include .env

COMPOSE_PROJECT_NAME ?= app
COMPOSE_FILE ?= ./docker/docker-compose.yml
COMPOSE_OPTIONS ?= -p $(COMPOSE_PROJECT_NAME) -f $(COMPOSE_FILE)

CONTAINER_GO ?= $(COMPOSE_PROJECT_NAME)_app_1
CONTAINER_DB ?= $(COMPOSE_PROJECT_NAME)_postgres_1

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
	@docker-compose $(COMPOSE_OPTIONS) up -d

stop:		## Stop application
	@docker-compose $(COMPOSE_OPTIONS) stop

down:		## Down application
	@docker-compose $(COMPOSE_OPTIONS) down

rebuild:	## Rebuild docker
	@docker-compose $(COMPOSE_OPTIONS) build --no-cache


db:		## Enter to DB
	@docker-compose $(COMPOSE_OPTIONS) \
		exec db \
		psql -U $(DB_USERNAME) $(DB_DATABASE)


tests:	## Run tests
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm app \
		go test -v -vet=off ./... 


migrate-create:	## Create migration [-n=name]
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm migrate \
		create -dir ./ -ext sql -seq $(n)

migrate-version: ## Get migration version
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ version 

migrate-up:
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ up 

migrate-down:
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ down 

migrate-force:	## Force migration [-v=version]
	@docker-compose $(COMPOSE_OPTIONS) \
		run --rm migrate \
		-database $(MIGRATE_DB_URL) -path ./ force $(v)
