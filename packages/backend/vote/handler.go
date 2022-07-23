package vote

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoteHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type voteHandler struct {
	service VoteService
}

func NewVoteHandler(service VoteService) VoteHandler {
	return &voteHandler{service}
}

func (h *voteHandler) getAllVotes(c *gin.Context) {
	// todo: by feedback
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllVotes not implemented",
	})
}

func (h *voteHandler) addVote(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "addVote not implemented",
	})
}

func (h *voteHandler) deleteVote(c *gin.Context) {
	// todo: when feedback is deleted -> delete related votes
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteVote not implemented",
	})
}
