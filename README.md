# Product Feedback App Backend

[![CI](https://github.com/Arthur199212/product-feedback-app/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/Arthur199212/product-feedback-app/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/Arthur199212/product-feedback-app/branch/master/graph/badge.svg?token=KPT5HEGLCH)](https://codecov.io/gh/Arthur199212/product-feedback-app)

## Technologies

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- JWT
- [PostgreSQL](https://www.postgresql.org/)
- [logrus](https://github.com/sirupsen/logrus)
- [Swagger](https://goswagger.io/)
- Websocket with [gorilla/websocket](https://github.com/gorilla/websocket)
- [golang/mock](https://github.com/golang/mock)

## How to Start

1. Run `make up` to spin up PostgreSQL DB with Docker compose.
1. Install [golang-migrate/migrate](https://github.com/golang-migrate/migrate) and run `make migrate_up` to run migrations.
1. Create `.env` file (see `.env.example` as an example).
1. Run `make run` to start server.

## Done

- [x] Authentication with GitHub & JWT
- [x] Create, Read operations with Users
- [x] Create, Read, Update, Delete operations with Feedback
- [x] Create, Read operations with Comments
- [x] Create, Read, Delete a Votes
- [x] Deployed (Heroku), [endpoint](https://go-product-feedback.herokuapp.com/)
- [x] Add [Swagger documentation](https://go-product-feedback.herokuapp.com/docs)
- [x] Setup CICD
- [x] Use Websocket to notify users about updates of feedback, comments & votes

## Todo

- [ ] Add unit-tests
- [ ] Add api-tests
- [ ] Update, Delete operations with Comments
- [ ] Update, Delete operations with Users

## How to run locally

1. Install Docker. Here is an example how it can be done for Windows [link](https://docs.docker.com/desktop/install/windows-install/)
1. Verify that it's installed and workes `docker version`.
1. Run `docker volume create product-feedback-db` to create a volume for PostgreSQL database to persist data.
1. Verify that it's created `docker volume ls` (should be seen in the output).
1. Add `.env` file to the root of the project. See `.env.example` for reference.
1. Run `docker compose up -d` from the root of the project.
1. App should start and be available on http://localhost:8000. Also, you may check Swagger documentation on http://localhost:8000/docs.
1. Run `docker compose down` to clean up.
