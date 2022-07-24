package vote

import (
	"database/sql"
	"fmt"
	"time"
)

type VoteRepository interface {
	Create(userId int, v createVoteInput) (int, error)
}

type voteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *voteRepository {
	return &voteRepository{db}
}

const (
	votesTable = "votes"
)

func (r *voteRepository) Create(userId int, v createVoteInput) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (feedback_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, votesTable)

	var id int
	currentTime := time.Now().UTC()
	err := r.db.QueryRow(
		query,
		v.FeedbackId,
		userId,
		currentTime,
		currentTime,
	).Scan(&id)

	return id, err
}
