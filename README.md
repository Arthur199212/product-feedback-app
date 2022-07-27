# Product Feedback App Backend

## Technologies

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- JWT
- [PostgreSQL](https://www.postgresql.org/)
- [logrus](https://github.com/sirupsen/logrus)
- [Heroku](https://www.heroku.com/)

## How to Start

1. Run `make up` to spin up PostgreSQL DB with Docker compose.
1. Install [golang-migrate/migrate](https://github.com/golang-migrate/migrate) and run `make migrate-up` to run migrations.
1. Create `.env` file (see `.env.example` as an example).
1. Run `make run` to start server.

## Done

- [x] Authentication with GitHub & JWT
- [x] Create, Read operations with Users
- [x] Create, Read, Update, Delete operations with Feedback
- [x] Create, Read operations with Comments
- [x] Create, Read, Delete a Votes
- [x] Deployed (Heroku), [endpoint](https://go-product-feedback.herokuapp.com/)

## Todo

- [ ] Setup CICD
- [ ] Add Swagger
- [ ] Add unit-tests
- [ ] Add api-tests
- [ ] Update, Delete operations with Comments
- [ ] Update, Delete operations with Users
