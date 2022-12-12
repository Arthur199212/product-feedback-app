package vote

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type VoteRepository interface {
	Create(userId int, v createVoteInput) (int, error)
	Delete(userId, voteId int) error
	GetAll(feedbackIds []int) ([]Vote, error)
	GetByFeedbackId(userId, feedbackId int) (Vote, error)
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

func (r *voteRepository) Delete(userId, voteId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1 AND id=$2
	`, votesTable)

	res, err := r.db.Exec(query, userId, voteId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err == nil && rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *voteRepository) GetAll(feedbackIds []int) ([]Vote, error) {
	whereClauseValues := make([]string, 0, 1)
	args := make([]interface{}, 0, 1)
	argId := 1

	if len(feedbackIds) > 0 {
		whereClauseValues = append(
			whereClauseValues,
			fmt.Sprintf("feedback_id = ANY($%d::int[])", argId),
		)
		args = append(args, pq.Array(feedbackIds))
		argId++
	}

	whereClauseQuery := strings.Join(whereClauseValues, " AND ")
	if len(whereClauseQuery) != 0 {
		whereClauseQuery = "WHERE " + whereClauseQuery
	}

	query := fmt.Sprintf(`
		SELECT id, feedback_id, user_id, created_at, updated_at FROM %s
		%s
	`, votesTable, whereClauseQuery)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	votes := []Vote{}
	for rows.Next() {
		var c Vote
		err = rows.Scan(
			&c.Id,
			&c.FeedbackId,
			&c.UserId,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return votes, err
		}
		votes = append(votes, c)
	}

	return votes, nil
}

func (r *voteRepository) GetByFeedbackId(
	userId,
	feedbackId int,
) (Vote, error) {
	query := fmt.Sprintf(`
		SELECT id, feedback_id, user_id, created_at, updated_at FROM %s
		WHERE user_id=$1 AND feedback_id=$2
	`, votesTable)

	var v Vote
	err := r.db.QueryRow(query, userId, feedbackId).Scan(
		&v.Id,
		&v.FeedbackId,
		&v.UserId,
		&v.CreatedAt,
		&v.UpdatedAt,
	)

	return v, err
}
