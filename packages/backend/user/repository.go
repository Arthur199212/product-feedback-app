package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserRepository interface {
	Create(user User) (int, error)
	GetByEmail(email string) (User, error)
	GetById(id int) (User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

const (
	UsersTable = "users"
)

func (r *userRepository) Create(user User) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (email, name, user_name, avatar_url, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6) returning id
	`, UsersTable)
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

func (r *userRepository) GetByEmail(email string) (User, error) {
	// todo
	return User{}, errors.New("not implemented")
}

func (r *userRepository) GetById(id int) (User, error) {
	// todo
	return User{}, errors.New("not implemented")
}
