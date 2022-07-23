package feedback

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type FeedbackRepository interface {
	Create(userId int, f createFeedbackInput) (int, error)
	Delete(userId, feedbackId int) error
	GetAll() ([]Feedback, error)
	GetById(userId, feedbackId int) (Feedback, error)
	Update(userId, feedbackId int, f updateFeedbackInput) error
}

type feedbackRepository struct {
	db *sql.DB
}

func NewFeedbackRepository(db *sql.DB) *feedbackRepository {
	return &feedbackRepository{db}
}

const (
	feedbackTable = "feedback"

	defaultFeedbackStatus = "idea"
)

var errNoInputToUpdate = errors.New("no input to update")

func (r *feedbackRepository) Create(userId int, f createFeedbackInput) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (title, body, category, status, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`, feedbackTable)

	var feedbackId int
	currentTime := time.Now().UTC()
	status := defaultFeedbackStatus
	if f.Status != nil {
		status = *f.Status
	}

	err := r.db.QueryRow(
		query,
		f.Title,
		f.Body,
		f.Category,
		status,
		userId,
		currentTime,
		currentTime,
	).Scan(&feedbackId)

	return feedbackId, err
}

func (r *feedbackRepository) Delete(userId, feedbackId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1 AND id=$2
	`, feedbackTable)

	res, err := r.db.Exec(query, userId, feedbackId)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err == nil && rows == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (r *feedbackRepository) GetAll() ([]Feedback, error) {
	// ORDER BY id DESC - shows later created first
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

func (r *feedbackRepository) Update(
	userId,
	feedbackId int,
	f updateFeedbackInput,
) error {
	setValues := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)
	argsId := 1

	if f.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argsId))
		args = append(args, f.Body)
		argsId++
	}

	if f.Category != nil {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argsId))
		args = append(args, f.Category)
		argsId++
	}

	if f.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argsId))
		args = append(args, f.Status)
		argsId++
	}

	if f.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsId))
		args = append(args, f.Title)
		argsId++
	}

	if len(setValues) == 0 {
		return errNoInputToUpdate
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`
		UPDATE %s SET %s
		WHERE user_id=$%d AND id=$%d
	`, feedbackTable, setQuery, argsId, argsId+1)
	args = append(args, userId, feedbackId)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err == nil && rows == 0 {
		return sql.ErrNoRows
	}

	return err
}
