postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine
createdb:
	docker exec -it postgres13 createdb --username=postgres --owner=postgres transfer-me
dropdb:
	docker exec -it postgres13 dropdb --username=postgres transfer-me
migrateup:
	migrate -path app/gateways/db/postgres/migrations -database "postgresql://postgres:postgres@localhost:5432/transfer-me?sslmode=disable" -verbose up
migratedown:
	migrate -path app/gateways/db/postgres/migrations -database "postgresql://postgres:postgres@localhost:5432/transfer-me?sslmode=disable" -verbose down
test:
	go test -v ./...
build:
	go build cmd/main.go
run:
	./run.sh
build-image:
	@echo "==> Building Docker API image"
	docker build . -t transfer-me

.PHONY: postgres createdb dropdb migrateup migratedown test build run build-image