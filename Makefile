include app.env

# Docker configuration
DOCKER_COMPOSE_FILE := docker-compose.yaml
SERVICE_NAME := api

.PHONY: up down logs
up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

.PHONY: createdb dropdb postgres postgresnetwork
postgresnetwork:
	docker network create -d bridge $(CONTAINER_NETWORK_NAME)

postgres:
	docker run --name $(CONTAINER_NAME) --network $(CONTAINER_NETWORK_NAME) -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:16-alpine

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)

.PHONY: migrateup migrateup1 migratedown migratedown1
migrateup:
	migrate -path sqlc/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migrateup1:
	migrate -path sqlc/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up 1

migratedown:
	migrate -path sqlc/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

migratedown1:
	migrate -path sqlc/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down 1

.PHONY: sqlc server test mock
sqlc:
	sqlc generate

server:
	go run cmd/main.go

test:
	go test -v -cover ./...

mock:
	mockgen -package mock -destination internal/store/mock/store.go your-project-name/internal/store Store

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir sqlc/migrations -seq $(name)