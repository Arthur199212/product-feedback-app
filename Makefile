# include .env
# export

MIGRATE_CMD=migrate -path ./migrations -database ${DATABASE_URL}

run:
	go run cmd/main.go

clean:
	rm bin -r || true

build: clean
	go build -v -o bin/main ./cmd

up:
	docker compose up -d

down:
	docker compose down

migrate_up:
	${MIGRATE_CMD} up

migrate_down:
	${MIGRATE_CMD} down

lint:
	gofmt -d .

lint_fix:
	gofmt -w .

check_swagger_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_install
	swagger generate spec -o ./swagger.yml --scan-models
