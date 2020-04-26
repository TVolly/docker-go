#!make
-include .env

COMPOSE_PROJECT_NAME ?= app
COMPOSE_FILE ?= ./docker/docker-compose.yml
COMPOSE_OPTIONS ?= -p $(COMPOSE_PROJECT_NAME) -f $(COMPOSE_FILE)

CONTAINER_DB ?= $(COMPOSE_PROJECT_NAME)_postgres_1

DB_DRIVER ?= postgres
DB_USERNAME ?= postgres
DB_PASSWORD ?= secret
DB_HOST ?= postgres
DB_PORT ?= 5432
DB_DATABASE ?= $(DB_USERNAME)

MIGRATE_DIR = /migrations/
MIGRATE_DB_URL = $(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable
MIGRATE_OPTIONS = --rm -v $(PWD)/migrations:/migrations --network=container:$(CONTAINER_DB) migrate/migrate


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
	@docker exec -it $(CONTAINER_DB) psql -U $(DB_USERNAME) $(DB_DATABASE)

migrate-create:	## Create migration [-n=name]
	@docker run $(MIGRATE_OPTIONS) create -dir $(MIGRATE_DIR) -ext sql -seq $(n)

migrate-version:	## Get migration version
	@docker run $(MIGRATE_OPTIONS) -database $(MIGRATE_DB_URL) -path $(MIGRATE_DIR) version

migrate-up:	## Run migration
	@docker run $(MIGRATE_OPTIONS) -database $(MIGRATE_DB_URL) -path $(MIGRATE_DIR) up

migrate-down:	## Rollback migration
	@docker run -it $(MIGRATE_OPTIONS) -database $(MIGRATE_DB_URL) -path $(MIGRATE_DIR) down

migrate-force:	## Force migration [-v=version]
	@docker run $(MIGRATE_OPTIONS) -database $(MIGRATE_DB_URL) -path $(MIGRATE_DIR) force $(v)
