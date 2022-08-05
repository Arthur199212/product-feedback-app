// Package classification Product Feedback API
//
// Documentation Product Feedback API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
//  SecurityDefinitions:
//  Bearer:
//    type: apiKey
//    in: header
//    name: Authorization
//
// swagger:meta
package provider

import (
	"net/http"
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	"product-feedback/middleware"
	"product-feedback/notifier"
	"product-feedback/user"
	"product-feedback/validation"
	"product-feedback/vote"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	go_openapi_middleware "github.com/go-openapi/runtime/middleware"
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
	Notifier notifier.NotifierHandler
}

func NewHandler(
	l *logrus.Logger,
	v *validation.Validation,
	s *Service,
) Handler {
	return &handler{
		Auth:     auth.NewAuthHandler(l, v, s.Auth),
		Comment:  comment.NewCommentHandler(l, v, s.Comment),
		Feedback: feedback.NewFeedbackHandler(l, v, s.Feedback),
		User:     user.NewUserHandler(l, v, s.User),
		Vote:     vote.NewVoteHandler(l, v, s.Vote),
		Notifier: notifier.NewNotifierHandler(l, &s.Notifier),
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(middleware.CorsMiddleware)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Product Feedback Api")
	})

	opts := go_openapi_middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := go_openapi_middleware.Redoc(opts, nil)

	router.StaticFile("swagger.yml", "./swagger.yml")
	router.GET("/docs", gin.WrapH(sh))

	api := router.Group("/api")

	h.Auth.AddRoutes(api)
	h.Comment.AddRoutes(api)
	h.Feedback.AddRoutes(api)
	h.User.AddRoutes(api)
	h.Vote.AddRoutes(api)
	h.Notifier.AddRoutes(api)

	return router
}
