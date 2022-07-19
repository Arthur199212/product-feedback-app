package comment

import "database/sql"

type CommentRepository interface {
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db}
}
