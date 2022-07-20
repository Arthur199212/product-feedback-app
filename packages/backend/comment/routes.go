package comment

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *commentHandler) AddRoutes(rg *gin.RouterGroup) {
	comments := rg.Group("/comments", middleware.AuthRequired)
	{
		comments.GET("/", h.getAllComments)
		comments.POST("/", h.addComment)
		comments.PUT("/:id", h.updateComment)
		comments.DELETE("/:id", h.deleteComment)
	}
}
