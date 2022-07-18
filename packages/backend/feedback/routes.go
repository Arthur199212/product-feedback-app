package feedback

import (
	"github.com/gin-gonic/gin"
)

func (h *FeedbackHandler) AddRoutes(rg *gin.RouterGroup) {
	feedback := rg.Group("/feedback")
	{
		feedback.GET("/", h.getAllFeedback)
		feedback.POST("/", h.createFeedback)
		feedback.GET("/:id", h.getFeedbackById)
		feedback.PUT("/:id", h.updateFeedback)
		feedback.DELETE("/:id", h.deleteFeedback)
	}
}
