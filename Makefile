.PHONY: env-setup env-teardown dev install-dependencies test

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

env-teardown:
	docker-compose -f docker-compose.dev.yml down

dev: env-setup
	air -c cmd/api/.air.toml

install-dependencies:
	go get github.com/cosmtrek/air@v1.15.1
	go mod tidy

setup-test-db:
	docker-compose -f docker-compose.test.yml up -d

teardown-test-db:
	docker-compose -f docker-compose.test.yml down

test:
	make setup-test-db
	go test -v -p 1 -count=1 ./...
	make teardown-test-db
