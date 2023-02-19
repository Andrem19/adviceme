migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root adviceme_db

dropdb:
	docker exec -it postgres12 dropdb adviceme_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/adviceme_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/adviceme_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/adviceme_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/adviceme_db?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: postgres createdb dropdb migratedown migrateup migratedown1 migrateup1 sqlc server test migrateinit