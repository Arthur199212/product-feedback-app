package vote

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	repo := NewVoteRepository()
	service := NewVoteService(repo)
	handler := NewVoteHandler(service)

	votes := rg.Group("/votes")
	{
		votes.GET("/", handler.getAllVotes)
		votes.POST("/", handler.addVote)
		votes.DELETE("/:id", handler.deleteVote)
	}
}
