package provider

import (
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	users "product-feedback/user"
	"product-feedback/vote"
)

type Service struct {
	Auth     auth.AuthService
	Comment  comment.CommentService
	Feedback feedback.FeedbackService
	User     users.UserService
	Vote     vote.VoteService
}

func NewService(r *Repository) *Service {
	return &Service{
		Auth:     auth.NewAuthService(r.User),
		Comment:  comment.NewCommentService(r.Comment),
		Feedback: feedback.NewFeedbackService(r.Feedback),
		User:     users.NewUserService(r.User),
		Vote:     vote.NewVoteService(r.Vote),
	}
}
