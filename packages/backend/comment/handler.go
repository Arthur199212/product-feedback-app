package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *CommentService
}

func NewCommentHandler(service *CommentService) *CommentHandler {
	return &CommentHandler{service}
}

func (h *CommentHandler) getAllComments(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllComments not implemented",
	})
}

func (h *CommentHandler) addComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "addComment not implemented",
	})
}

func (h *CommentHandler) updateComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateComment not implemented",
	})
}

func (h *CommentHandler) deleteComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteComment not implemented",
	})
}
