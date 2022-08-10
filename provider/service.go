package provider

import (
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	"product-feedback/notifier"
	"product-feedback/user"
	"product-feedback/vote"
)

type Service struct {
	Auth     auth.AuthService
	Comment  comment.CommentService
	Feedback feedback.FeedbackService
	User     user.UserService
	Vote     vote.VoteService
	Notifier notifier.NotifierService
}

func NewService(r *Repository) *Service {
	notifierService := notifier.NewNotifierSerivice()
	userService := user.NewUserService(r.User)

	return &Service{
		Auth:     auth.NewAuthService(userService),
		Comment:  comment.NewCommentService(r.Comment, notifierService),
		Feedback: feedback.NewFeedbackService(r.Feedback, notifierService),
		User:     userService,
		Vote:     vote.NewVoteService(r.Vote, notifierService),
		Notifier: notifierService,
	}
}
