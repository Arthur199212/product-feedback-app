package comment

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	repo := NewCommentRepository()
	service := NewCommentService(repo)
	handler := NewCommentHandler(service)

	comments := rg.Group("/comments")
	{
		comments.GET("/", handler.getAllComments)
		comments.POST("/", handler.addComment)
		comments.PUT("/:id", handler.updateComment)
		comments.DELETE("/:id", handler.deleteComment)
	}
}
