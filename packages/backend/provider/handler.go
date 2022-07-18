package provider

import (
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	users "product-feedback/user"
	"product-feedback/vote"
)

type Handler struct {
	Authorization *auth.AuthHandler
	Comment       *comment.CommentHandler
	Feedback      *feedback.FeedbackHandler
	User          *users.UserHandler
	Vote          *vote.VoteHandler
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		Authorization: auth.NewAuthHandler(s.Authorization),
		Comment:       comment.NewCommentHandler(s.Comment),
		Feedback:      feedback.NewFeedbackHandler(s.Feedback),
		User:          users.NewUserHandler(s.User),
		Vote:          vote.NewVoteHandler(s.Vote),
	}
}
