package provider

import (
	"database/sql"
	"product-feedback/comment"
	"product-feedback/feedback"
	"product-feedback/user"
	"product-feedback/vote"
)

type Repository struct {
	Comment  comment.CommentRepository
	Feedback feedback.FeedbackRepository
	User     user.UserRepository
	Vote     vote.VoteRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Comment:  comment.NewCommentRepository(db),
		Feedback: feedback.NewFeedbackRepository(db),
		User:     user.NewUserRepository(db),
		Vote:     vote.NewVoteRepository(db),
	}
}
