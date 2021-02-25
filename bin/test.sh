#!/bin/bash

set -e

goose -dir database/migrations -table "migration_versions" postgres "$DATABASE_URL" up
DATABASE_URL=$DATABASE_URL go test -v -p 1 -count=1 ./...
