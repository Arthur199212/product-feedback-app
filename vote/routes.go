package vote

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *voteHandler) AddRoutes(rg *gin.RouterGroup) {
	votes := rg.Group("/votes", middleware.AuthRequired)
	{
		votes.GET("/", h.getAllVotes)
		votes.POST("/", h.createVote)
		votes.DELETE("/:id", h.deleteVote)
		votes.POST("/toggle", h.toggleVote)
	}
}
