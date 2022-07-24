package vote

import (
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VoteHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type voteHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service VoteService
}

func NewVoteHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service VoteService,
) VoteHandler {
	return &voteHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

func (h *voteHandler) getAllVotes(c *gin.Context) {
	// todo: by feedback
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllVotes not implemented",
	})
}

type createVoteInput struct {
	FeedbackId int `json:"feedbackId" validate:"required,gt=0"`
}

func (h *voteHandler) createVote(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	var input createVoteInput
	if err = c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Vote is invalid",
		})
		return
	}

	if err = h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Vote is invalid",
		})
		return
	}

	voteId, err := h.service.Create(userId, input)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"voteId": voteId,
	})
}

func (h *voteHandler) deleteVote(c *gin.Context) {
	// todo: when feedback is deleted -> delete related votes
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteVote not implemented",
	})
}
