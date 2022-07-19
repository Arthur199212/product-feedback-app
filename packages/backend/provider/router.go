package provider

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Product Feedback Api")
	})

	api := router.Group("/api")
	h.Authorization.AddRoutes(api)
	h.Comment.AddRoutes(api)
	h.Feedback.AddRoutes(api)
	h.User.AddRoutes(api)
	h.Vote.AddRoutes(api)

	return router
}
