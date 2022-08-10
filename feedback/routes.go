package feedback

import "github.com/gin-gonic/gin"

func (h *feedbackHandler) AddRoutes(
	rg *gin.RouterGroup,
	authMiddleware gin.HandlerFunc,
) {
	feedback := rg.Group("/feedback", authMiddleware)
	{
		feedback.GET("/", h.GetAllFeedback)
		feedback.POST("/", h.CreateFeedback)
		feedback.GET("/:id", h.GetFeedbackById)
		feedback.PUT("/:id", h.UpdateFeedback)
		feedback.DELETE("/:id", h.DeleteFeedback)
	}
}
