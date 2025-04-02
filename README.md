## Development

## Go Install

Download from [go lang](https://go.dev/doc/install).

After save global path
```bash
echo 'export GOROOT=/usr/local/go' >> ~/.zshrc
echo 'export PATH=$GOROOT/bin:$PATH' >> ~/.zshrc
source ~/.zshrc
```

## Create Database
Create a database and user with privileges.
```bash
docker run --name my-mysql \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=gigbuddy \
  -e MYSQL_USER=admin \
  -e MYSQL_PASSWORD=admin \
  -p 3306:3306 \
  -d mysql:latest
```
#### Connect to database
```bash
docker exec -it mysql-server mysql -u root -p
```
## Migrations

We use [goose](https://github.com/pressly/goose) for migrations with MySQL.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Migrations
First step set environment variables:

```bash
export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING="root:root@tcp(localhost:3306)/gigbuddy?parseTime=true"
export GOOSE_MIGRATIONS_DIR=migrations
```

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

## Swagger - Api Documentation

Swagger docs are available at `http://localhost:8080/docs`
When you change the code you need to run the following command to update the docs:

```bash
swag init -g server.go
```

### Firebase
Please contact admin for firebase config files.

### Before Launch
For before launch script to work you need to make it executable.
```bash
chmod +x sh/before_launch.sh
```