postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=quynhnhu2010 -d postgres:12-alpine

restartpq:
	docker start postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:quynhnhu2010@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:quynhnhu2010@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratemyup:
	migrate -path db/migration -database "mysql://root:quynhnhu2010@localhost:3306/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate

commit:
	git add .
	git commit -m "automatically commit"
	git push origin main 

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/minhtam3010/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock