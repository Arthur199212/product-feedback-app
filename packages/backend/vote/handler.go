package vote

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	service *VoteService
}

func NewVoteHandler(service *VoteService) *VoteHandler {
	return &VoteHandler{service}
}

func (h *VoteHandler) getAllVotes(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllVotes not implemented",
	})
}

func (h *VoteHandler) addVote(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "addVote not implemented",
	})
}

func (h *VoteHandler) deleteVote(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteVote not implemented",
	})
}
