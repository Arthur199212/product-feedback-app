package provider

import (
	"net/http"
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	users "product-feedback/user"
	"product-feedback/vote"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	InitRoutes() *gin.Engine
}

type handler struct {
	Auth     auth.AuthHandler
	Comment  comment.CommentHandler
	Feedback feedback.FeedbackHandler
	User     users.UserHandler
	Vote     vote.VoteHandler
}

func NewHandler(s *Service) Handler {
	return &handler{
		Auth:     auth.NewAuthHandler(s.Auth),
		Comment:  comment.NewCommentHandler(s.Comment),
		Feedback: feedback.NewFeedbackHandler(s.Feedback),
		User:     users.NewUserHandler(s.User),
		Vote:     vote.NewVoteHandler(s.Vote),
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Product Feedback Api")
	})

	api := router.Group("/api")

	h.Auth.AddRoutes(api)
	h.Comment.AddRoutes(api)
	h.Feedback.AddRoutes(api)
	h.User.AddRoutes(api)
	h.Vote.AddRoutes(api)

	return router
}
