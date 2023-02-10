package user

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type UserRepository interface {
	Create(user User) (int, error)
	GetByEmail(email string) (User, error)
	GetById(id int) (User, error)
	GetAll(userIds []int) ([]User, error)
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

func (r *userRepositoryPostgres) GetAll(userIds []int) ([]User, error) {
	whereClauseValues := make([]string, 0, 1)
	args := make([]interface{}, 0, 1)
	argId := 1

	if len(userIds) > 0 {
		whereClauseValues = append(
			whereClauseValues,
			fmt.Sprintf("id = ANY($%d::int[])", argId),
		)
		args = append(args, pq.Array(userIds))
		argId++
	}

	whereClauseQuery := strings.Join(whereClauseValues, " AND ")
	if len(whereClauseQuery) != 0 {
		whereClauseQuery = "WHERE " + whereClauseQuery
	}

	// ORDER BY id ASC - shows earlier created first
	query := fmt.Sprintf(`
		SELECT id, email, name, user_name, avatar_url, created_at, updated_at FROM %s
		%s
		ORDER BY id ASC
	`, usersTable, whereClauseQuery)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var c User
		err = rows.Scan(
			&c.Id,
			&c.Email,
			&c.Name,
			&c.UserName,
			&c.AvatarUrl,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, c)
	}

	return users, nil
}
