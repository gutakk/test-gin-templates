#!/bin/bash

set -e

goose -dir database/migrations -table "migration_versions" postgres "postgres://postgres:postgres@test-gin-templates_db_test/test-gin-templates_test?sslmode=disable" up
DATABASE_URL=postgres://postgres:postgres@test-gin-templates_db_test/test-gin-templates_test?sslmode=disable go test -v -p 1 -count=1 ./...
