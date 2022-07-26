package comment

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *commentHandler) AddRoutes(rg *gin.RouterGroup) {
	comments := rg.Group("/comments", middleware.AuthRequired)
	{
		comments.GET("/", h.getAllComments)
		comments.GET("/:id", h.getCommentById)
		comments.POST("/", h.createComment)
		comments.PUT("/:id", h.updateComment)
		comments.DELETE("/:id", h.deleteComment)
	}
}
