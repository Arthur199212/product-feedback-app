package feedback

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	repo := NewFeedbackRepository()
	service := NewFeedbackService(repo)
	handler := NewFeedbackHandler(service)

	feedback := rg.Group("/feedback")
	{
		feedback.GET("/", handler.getAllFeedback)
		feedback.POST("/", handler.createFeedback)
		feedback.GET("/:id", handler.getFeedbackById)
		feedback.PUT("/:id", handler.updateFeedback)
		feedback.DELETE("/:id", handler.deleteFeedback)
	}
}
