#!/bin/sh
export GOOSE_DRIVER=mysql
export GOOSE_DBSTRING="root:root@tcp(localhost:3306)/gigbuddy?parseTime=true"
export GOOSE_MIGRATIONS_DIR=migrations
swag init -g server.go
