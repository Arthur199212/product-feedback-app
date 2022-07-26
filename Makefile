-include .env
export

MIGRATE_CMD=migrate -path ./migrations -database ${DATABASE_URL}
GOFILES_WITHOUT_VENDOR=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

run:
	go run cmd/main.go

clean:
	rm bin -r || true

build: clean
	go build -v -o bin/main ./cmd

test:
	go test -v ./...

test_coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./... && \
	go tool cover -html=coverage.out -o coverage.html

up:
	docker compose up -d

down:
	docker compose down

migrate_create:
	migrate create -ext sql -dir ./migrations -seq ${NAME}

migrate_up:
	${MIGRATE_CMD} up

migrate_down:
	${MIGRATE_CMD} down 1

lint:
	@if [ -n "$$(gofmt -l ${GOFILES_WITHOUT_VENDOR})" ]; \
		then echo 'Forgot to run "make lint_fix"?' && exit 1; \
	fi

lint_fix:
	@gofmt -l -w ${GOFILES_WITHOUT_VENDOR}

check_swagger_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_install
	swagger generate spec -o ./swagger.yml --scan-models
