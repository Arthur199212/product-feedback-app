package provider

import (
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	"product-feedback/user"
	"product-feedback/vote"
)

type Service struct {
	Auth     auth.AuthService
	Comment  comment.CommentService
	Feedback feedback.FeedbackService
	User     user.UserService
	Vote     vote.VoteService
}

func NewService(r *Repository) *Service {
	userService := user.NewUserService(r.User)

	return &Service{
		Auth:     auth.NewAuthService(userService),
		Comment:  comment.NewCommentService(r.Comment),
		Feedback: feedback.NewFeedbackService(r.Feedback),
		User:     userService,
		Vote:     vote.NewVoteService(r.Vote),
	}
}
