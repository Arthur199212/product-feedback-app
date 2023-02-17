package vote

import (
	"database/sql"
	"fmt"
	"net/http"
	"product-feedback/middleware"
	"product-feedback/utils"
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

// swagger:route GET /api/votes votes getAllVotes
// Returns a list of votes in the system
//
// security:
// - Bearer:
//
// responses:
//	200: getAllVotesResponse

func (h *voteHandler) getAllVotes(c *gin.Context) {
	feedbackIds, err := utils.ParseIdsFromQuery(c.Query("feedbackId"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	votes, err := h.service.GetAll(feedbackIds)
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
	// The id of the feedback vote is related to
	//
	// required: true
	// min: 1
	FeedbackId int `json:"feedbackId" validate:"required,gt=0"`
}

// swagger:route POST /api/votes votes createVote
// Creates a vote
//
// security:
// - Bearer:
//
// responses:
//	200: createVoteResponse

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

	voteId, err := h.service.Create(userId, input.FeedbackId)
	switch err {
	case nil:
		break
	case ErrVoteAlreadyExists:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Vote already exists",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"voteId": voteId,
	})
}

// swagger:route DELETE /api/votes/:id votes deleteVote
// Deletes a vote
//
// security:
// - Bearer:
//
// responses:
//	200: okResponse
//	404: errorResponse

func (h *voteHandler) deleteVote(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	voteIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "voteId is invalid",
		})
		return
	}

	if err = h.v.ValidateVar(voteIdInt, "required,gt=0"); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "voteId is invalid",
		})
		return
	}

	err = h.service.Delete(userId, voteIdInt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "Vote not found",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

type toggleVoteInput struct {
	// The id of the feedback vote is related to
	//
	// required: true
	// min: 1
	FeedbackId int `json:"feedbackId" validate:"required,gt=0"`
}

// swagger:route POST /api/votes/toggle votes toggleVote
// Toggles a vote
//
// security:
// - Bearer:
//
// responses:
//	200: okResponse
//	500: errorResponse

func (h *voteHandler) toggleVote(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	var input toggleVoteInput
	if err = c.BindJSON(&input); err != nil {
		h.l.Error(err)
		err = fmt.Errorf("invalid input: %+v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if err = h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		err = fmt.Errorf("invalid argument: %+v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	err = h.service.Toggle(userId, input.FeedbackId)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
