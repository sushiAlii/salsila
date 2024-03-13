include .env

dev:
	air

createdb:
	docker exec -it $(DOCKER_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it $(DOCKER_CONTAINER) dropdb $(DB_NAME) --username=$(DB_USER)

migrate-up:
	migrate -path pkg/db/migrations -database "$(DB_CONNECTION)" -verbose up

migrate-down:
	migrate -path pkg/db/migrations -database "$(DB_CONNECTION)" -verbose down

test:
	go test ./...

.PHONY: createdb dropdb migrate-up migrate-down dev test