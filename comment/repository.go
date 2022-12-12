package comment

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type CommentRepository interface {
	Create(userId int, f createCommentInput) (int, error)
	GetAll(feedbackIds []int) ([]Comment, error)
	GetById(commentId int) (Comment, error)
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
		INSERT INTO %s (body, feedback_id, user_id, parent_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`, commentsTable)
	var id int
	currentTime := time.Now().UTC()

	err := r.db.QueryRow(
		query,
		f.Body,
		f.FeedbackId,
		userId,
		f.ParentId,
		currentTime,
		currentTime,
	).Scan(&id)

	return id, err
}

func (r *commentRepository) GetAll(feedbackIds []int) ([]Comment, error) {
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

	// ORDER BY id ASC - shows earlier created first
	query := fmt.Sprintf(`
		SELECT id, body, feedback_id, user_id, parent_id, created_at, updated_at FROM %s
		%s
		ORDER BY id ASC
	`, commentsTable, whereClauseQuery)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		err = rows.Scan(
			&c.Id,
			&c.Body,
			&c.FeedbackId,
			&c.UserId,
			&c.ParentId,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *commentRepository) GetById(commentId int) (Comment, error) {
	query := fmt.Sprintf(`
		SELECT id, body, feedback_id, user_id, parent_id, created_at, updated_at FROM %s
		WHERE id=$1
	`, commentsTable)

	var c Comment
	err := r.db.QueryRow(query, commentId).Scan(
		&c.Id,
		&c.Body,
		&c.FeedbackId,
		&c.UserId,
		&c.ParentId,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	return c, err
}
