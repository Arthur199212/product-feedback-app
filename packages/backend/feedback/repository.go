package feedback

import (
	"database/sql"
	"fmt"
	"time"
)

type FeedbackRepository interface {
	Create(userId int, f createFeedbackInput) (int, error)
	GetAll() ([]Feedback, error)
	GetById(userId, feedbackId int) (Feedback, error)
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

func (r *feedbackRepository) GetAll() ([]Feedback, error) {
	// ORDER BY id DESC - shows latest created first
	query := fmt.Sprintf(`
		SELECT id, title, body, category, status, user_id, created_at, updated_at FROM %s
		ORDER BY id DESC
	`, feedbackTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fList := []Feedback{}
	for rows.Next() {
		var f Feedback
		err := rows.Scan(
			&f.Id,
			&f.Title,
			&f.Body,
			&f.Category,
			&f.Status,
			&f.UserId,
			&f.CreatedAt,
			&f.UpdatedAt,
		)
		if err != nil {
			return fList, err
		}
		fList = append(fList, f)
	}

	return fList, nil
}

func (r *feedbackRepository) GetById(userId, feedbackId int) (Feedback, error) {
	var f Feedback
	query := fmt.Sprintf(`
		SELECT id, title, body, category, status, user_id, created_at, updated_at FROM %s
		WHERE user_id=$1 AND id=$2
	`, feedbackTable)

	err := r.db.QueryRow(query, userId, feedbackId).Scan(
		&f.Id,
		&f.Title,
		&f.Body,
		&f.Category,
		&f.Status,
		&f.UserId,
		&f.CreatedAt,
		&f.UpdatedAt,
	)

	return f, err
}
