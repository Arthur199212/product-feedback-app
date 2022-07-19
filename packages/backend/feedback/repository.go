package feedback

import "database/sql"

type FeedbackRepository interface {
}

type feedbackRepository struct {
	db *sql.DB
}

func NewFeedbackRepository(db *sql.DB) *feedbackRepository {
	return &feedbackRepository{db}
}
