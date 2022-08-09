package feedback

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *feedbackHandler) AddRoutes(rg *gin.RouterGroup) {
	feedback := rg.Group("/feedback", middleware.AuthRequired)
	{
		feedback.GET("/", h.GetAllFeedback)
		feedback.POST("/", h.CreateFeedback)
		feedback.GET("/:id", h.GetFeedbackById)
		feedback.PUT("/:id", h.updateFeedback)
		feedback.DELETE("/:id", h.DeleteFeedback)
	}
}
