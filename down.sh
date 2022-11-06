#!/bin/sh

export DB_URL='postgres://postgres:postgres@localhost:5454/postgres?sslmode=disable'

migrate -database ${DB_URL} -path db/migrates down