package router

import (
	"product-feedback/auth"
	"product-feedback/comment"
	"product-feedback/feedback"
	users "product-feedback/user"
	"product-feedback/vote"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	auth.AddRoutes(api)
	comment.AddRoutes(api)
	feedback.AddRoutes(api)
	users.AddRoutes(api)
	vote.AddRoutes(api)

	return router
}
