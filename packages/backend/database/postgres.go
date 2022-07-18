package database

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DBName   string
	Host     string
	Password string
	Port     string
	SSLMode  string
	Username string
}

func NewPostgresDB(c Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.DBName, c.SSLMode,
	)
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
