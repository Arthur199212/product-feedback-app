package feedback

import (
	"database/sql"
	"fmt"
	"time"
)

type FeedbackRepository interface {
	Create(userId int, f createFeedbackInput) (int, error)
}

type feedbackRepository struct {
	db *sql.DB
}

func NewFeedbackRepository(db *sql.DB) *feedbackRepository {
	return &feedbackRepository{db}
}

const (
	feedbackTable = "feedback"
)

func (r *feedbackRepository) Create(userId int, f createFeedbackInput) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (title, body, category, status, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`, feedbackTable)
	var feedbackId int
	currentTime := time.Now().UTC()

	err := r.db.QueryRow(
		query,
		f.Title,
		f.Body,
		f.Category,
		f.Status,
		userId,
		currentTime,
		currentTime,
	).Scan(&feedbackId)

	return feedbackId, err
}
