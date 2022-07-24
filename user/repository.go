package user

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepository interface {
	Create(user User) (int, error)
	GetByEmail(email string) (User, error)
	GetById(id int) (User, error)
}

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{db}
}

const (
	usersTable = "users"
)

func (r *userRepositoryPostgres) Create(user User) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (email, name, user_name, avatar_url, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id
	`, usersTable)
	currentTime := time.Now().UTC()
	row := r.db.QueryRow(
		query,
		user.Email,
		user.Name,
		user.UserName,
		user.AvatarUrl,
		currentTime,
		currentTime,
	)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *userRepositoryPostgres) GetByEmail(email string) (User, error) {
	query := fmt.Sprintf(`
		SELECT id, email, name, user_name, avatar_url, created_at, updated_at FROM %s
		WHERE email=$1
	`, usersTable)

	var user User
	err := r.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.UserName,
		&user.AvatarUrl,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (r *userRepositoryPostgres) GetById(id int) (User, error) {
	query := fmt.Sprintf(`
		SELECT id, email, name, user_name, avatar_url, created_at, updated_at FROM %s
		WHERE id=$1
	`, usersTable)

	var user User
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.UserName,
		&user.AvatarUrl,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
