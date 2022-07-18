package comment

import (
	"github.com/gin-gonic/gin"
)

func (h *CommentHandler) AddRoutes(rg *gin.RouterGroup) {
	comments := rg.Group("/comments")
	{
		comments.GET("/", h.getAllComments)
		comments.POST("/", h.addComment)
		comments.PUT("/:id", h.updateComment)
		comments.DELETE("/:id", h.deleteComment)
	}
}
