include .env
export

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

migrate-up:
	${MIGRATE_CMD} up

migrate-down:
	${MIGRATE_CMD} down

lint:
	gofmt -d .

lint-fix:
	gofmt -w .
