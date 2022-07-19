package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type feedbackHandler struct {
	service FeedbackService
}

func NewFeedbackHandler(service FeedbackService) FeedbackHandler {
	return &feedbackHandler{service}
}

func (h *feedbackHandler) getAllFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllFeedback not implemented",
	})
}

func (h *feedbackHandler) createFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "createFeedback not implemented",
	})
}

func (h *feedbackHandler) getFeedbackById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getFeedbackById not implemented",
	})
}

func (h *feedbackHandler) updateFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateFeedback not implemented",
	})
}

func (h *feedbackHandler) deleteFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteFeedback not implemented",
	})
}
