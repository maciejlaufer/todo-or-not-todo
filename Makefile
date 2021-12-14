
-include app.env

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) todo_or_not_todo

dropdb:
	docker exec -it postgres14 dropdb todo_or_not_todo

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server