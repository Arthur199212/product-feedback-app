package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	service *FeedbackService
}

func NewFeedbackHandler(service *FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service}
}

func (h *FeedbackHandler) getAllFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllFeedback not implemented",
	})
}

func (h *FeedbackHandler) createFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "createFeedback not implemented",
	})
}

func (h *FeedbackHandler) getFeedbackById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getFeedbackById not implemented",
	})
}

func (h *FeedbackHandler) updateFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateFeedback not implemented",
	})
}

func (h *FeedbackHandler) deleteFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteFeedback not implemented",
	})
}
