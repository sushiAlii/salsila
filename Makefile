dev:
	air

createdb:
	docker exec -it 'rte-backend-db-1' createdb --username=postgres --owner=postgres salsila

dropdb:
	docker exec -it 'rte-backend-db-1' dropdb salsila --username=postgres

migrate-up:
	migrate -path pkg/db/migrations -database "postgresql://postgres:december181996@localhost:5432/salsila?sslmode=disable" -verbose up

migrate-down:
	migrate -path pkg/db/migrations -database "postgresql://postgres:december181996@localhost:5432/salsila?sslmode=disable" -verbose down

.PHONY: createdb dropdb migrate-up migrate-down dev