MIGRATE_CMD=migrate -path ./migrations -database 'postgres://postgres:secret@localhost:8010/postgres?sslmode=disable'

run:
	go run cmd/main.go

clean:
	rm bin -r || true

build: clean
	go build -v -o bin/main.exe ./cmd

up:
	docker compose up -d

down:
	docker compose down

migrate-up:
	${MIGRATE_CMD} up

migrate-down:
	${MIGRATE_CMD} down
