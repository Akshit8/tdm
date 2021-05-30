fmt:
	@echo "formatting code"
	go fmt ./...

lint:
	@echo "Linting source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

test:
	@echo "running tests"
	go test ./...

generate:
	go generate ./...

install:
	@echo "installing external dependencies"
	go mod download

graphql:
	@echo "generating graphql stubs"
	go run github.com/99designs/gqlgen generate

sqlc:

createdb:
	docker exec -it tdm_postgres_1 createdb --username=root --owner=root tdm

dropdb:
	docker exec -it tdm_postgres_1 dropdb tdm

migrationup:
	migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/tdm?sslmode=disable" -verbose up

migrationdown:
	migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/tdm?sslmode=disable" -verbose down

run:
	go run cmd/rest-server/main.go --env .dev.env

live:
	reflex -r '\.go' -s -- sh -c "make run"

dev:
	docker-compose up -d