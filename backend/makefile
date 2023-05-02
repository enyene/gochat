postgres:
	docker run --name userpostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

psql:
	docker exec -it userpostgres psql

createdb:
	docker exec -it userpostgres createdb --username=root --owner=root gochat

dropdb:
	docker exec -it userpostgres dropdb gochat

migrateup:
	migrate -path db/migrations -databse "postgresql://root:password@localhost:5432/gochat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -databse "postgresql://root:password@localhost:5432/gochat?sslmode=disable" -verbose up

migratedb:
	migrate create -ext sql -dir db/migrations add_user_table

.PHONY: postgres psql createdb dropdb migratedown migrateup migratedb