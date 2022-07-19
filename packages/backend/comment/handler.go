package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type commentHandler struct {
	service CommentService
}

func NewCommentHandler(service CommentService) CommentHandler {
	return &commentHandler{service}
}

func (h *commentHandler) getAllComments(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllComments not implemented",
	})
}

func (h *commentHandler) addComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "addComment not implemented",
	})
}

func (h *commentHandler) updateComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateComment not implemented",
	})
}

func (h *commentHandler) deleteComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteComment not implemented",
	})
}
