.PHONY: postgres createdb dropdb migrateup migratedown

postgres:
	docker run --name postgress15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgress15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path changelog -database "postgresql://user:password@localhost:5432/db2?sslmode=disable" -verbose up

migrateup_rm_test:
	rm ./changelog/000100_test_data.up.sql ./changelog/000200_test_data.down.sql

migrateup_remote:
	migrate -path changelog -database "postgresql://user:password@176.113.83.101:5432/db2?sslmode=disable" -verbose up

migratedown:
	migrate -path changelog -database "postgresql://user:password@localhost:5432/db2?sslmode=disable" -verbose down

migratedown_remote:
	migrate -path changelog -database "postgresql://user:password@176.113.83.101:5432/db?sslmode=disable" -verbose down