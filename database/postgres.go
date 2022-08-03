package database

import (
	"database/sql"
)

			func NewPostgresDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		connStr,
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
