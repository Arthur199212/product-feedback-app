FROM golang:1.19.2-alpine3.16 AS builder

WORKDIR /app

COPY . .

RUN go build -v -o main ./cmd

# Install golang-migrate
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY .env .
COPY migrations ./migrations
COPY wait-for.sh .
COPY start.sh .
COPY swagger.yml .

EXPOSE 8000

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
