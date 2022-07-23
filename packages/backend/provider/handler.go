package provider

import (
	"net/http"
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	"product-feedback/middleware"
	"product-feedback/user"
	"product-feedback/validation"
	"product-feedback/vote"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	InitRoutes() *gin.Engine
}

type handler struct {
	Auth     auth.AuthHandler
	Comment  comment.CommentHandler
	Feedback feedback.FeedbackHandler
	User     user.UserHandler
	Vote     vote.VoteHandler
}

func NewHandler(
	l *logrus.Logger,
	v *validation.Validation,
	s *Service,
) Handler {
	return &handler{
		Auth:     auth.NewAuthHandler(l, s.Auth),
		Comment:  comment.NewCommentHandler(l, v, s.Comment),
		Feedback: feedback.NewFeedbackHandler(l, v, s.Feedback),
		User:     user.NewUserHandler(l, v, s.User),
		Vote:     vote.NewVoteHandler(s.Vote),
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(middleware.CorsMiddleware)

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
