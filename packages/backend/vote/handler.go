package vote

import (
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"
	"strconv"

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
	var feedbackIdInt *int
	if feedbackId := c.Query("feedbackId"); feedbackId != "" {
		parsedFeedbackId, err := strconv.Atoi(feedbackId)
		if err != nil {
			h.l.Warn(err)
		} else {
			feedbackIdInt = &parsedFeedbackId
		}
	}

	if err := h.v.ValidateVar(feedbackIdInt, "omitempty,gt=0"); err != nil {
		h.l.Warn(err)
		feedbackIdInt = nil
	}

	votes, err := h.service.GetAll(feedbackIdInt)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, votes)
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
