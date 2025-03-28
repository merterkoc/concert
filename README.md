## Development
### Migrations


### Setup

We use [goose](https://github.com/pressly/goose) for migrations with MySQL.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

First step set environment variables:

```bash
export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING="root:root@tcp(localhost:3306)/gigbuddy?parseTime=true"
export GOOSE_MIGRATIONS_DIR=migrations

Create a new migration:

```bash
goose -s -dir migrations create <your_migration_name> sql
```

Run migrations:
```bash
goose -s -dir migrations up
```

Check migrations:
```bash
goose -s -dir migrations status
```

### Swagger 

Swagger docs are available at `http://localhost:8080/docs`
When you change the code you need to run the following command to update the docs:

```bash
swag init -g server.go
```