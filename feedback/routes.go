package feedback

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *feedbackHandler) AddRoutes(rg *gin.RouterGroup) {
	feedback := rg.Group("/feedback", middleware.AuthRequired)
	{
		feedback.GET("/", h.getAllFeedback)
		feedback.POST("/", h.createFeedback)
		feedback.GET("/:id", h.getFeedbackById)
		feedback.PUT("/:id", h.updateFeedback)
		feedback.DELETE("/:id", h.deleteFeedback)
	}
}
