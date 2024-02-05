postgres:
	docker run --name praktek -p 5434:5432 -e POSTGRES_PASSWORD=cuankipintar -d postgres:16-alpine

createdb:
	docker exec -it praktek createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it praktek dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:cuankipintar@localhost:5434/simple_bank?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://postgres:cuankipintar@localhost:5434/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test