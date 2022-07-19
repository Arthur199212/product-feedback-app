package vote

import (
	"github.com/gin-gonic/gin"
)

func (h *voteHandler) AddRoutes(rg *gin.RouterGroup) {
	votes := rg.Group("/votes")
	{
		votes.GET("/", h.getAllVotes)
		votes.POST("/", h.addVote)
		votes.DELETE("/:id", h.deleteVote)
	}
}
