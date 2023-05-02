postgres:
	docker run --name userpostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

psql:
	docker exec -it userpostgres psql

createdb:
	docker exec -it userpostgres createdb --username=root --owner=root gochat

dropdb:
	docker exec -it userpostgres dropdb gochat


.PHONY: postgres psql createdb dropdb