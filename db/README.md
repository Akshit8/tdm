## Setting Postgres with Docker

```bash
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine
```

Or use docker-compose.

## Migration tool

[migrate](https://github.com/golang-migrate/migrate)

## Migrations

Create:

```bash
migrate create -ext sql -dir db/migrations/ <migration name>
```

Up:

```bash
migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/tdm?sslmode=disable" -verbose up
```

Down:

```bash
migrate -path ./db/migrations -database "postgres://root:secret@localhost:5432/tdm?sslmode=disable" -verbose down
```