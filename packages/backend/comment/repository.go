package comment

import (
	"database/sql"
	"fmt"
	"time"
)

type CommentRepository interface {
	Create(userId int, f createCommentInput) (int, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db}
}

const (
	commentsTable = "comments"
)

func (r *commentRepository) Create(userId int, f createCommentInput) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (body, feedback_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`, commentsTable)
	var id int
	currentTime := time.Now().UTC()

	err := r.db.QueryRow(
		query,
		f.Body,
		f.FeedbackId,
		userId,
		currentTime,
		currentTime,
	).Scan(&id)

	return id, err
}
