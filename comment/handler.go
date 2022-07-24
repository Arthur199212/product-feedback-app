package comment

import (
	"database/sql"
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CommentHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type commentHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service CommentService
}

func NewCommentHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service CommentService,
) CommentHandler {
	return &commentHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

func (h *commentHandler) getAllComments(c *gin.Context) {
	// todo: implement options:
	// filter by: userId
	// sorted: date of creation, date of update
	// pagination: limit/size=<uint>, page=<uint>

	var feedbackIdInt *int
	if feedbackId := c.Query("feedbackId"); feedbackId != "" {
		parsedFeedbackId, err := strconv.Atoi(feedbackId)
		if err != nil {
			h.l.Warn(err)
		} else {
			feedbackIdInt = &parsedFeedbackId
		}
	}

	comments, err := h.service.GetAll(feedbackIdInt)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

type createCommentInput struct {
	Body       string `json:"body" validate:"required,min=5,max=255"`
	FeedbackId int    `json:"feedbackId" validate:"required,gt=0"`
}

func (h *commentHandler) createComment(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	var input createCommentInput
	if err = c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Input is invalid",
		})
		return
	}

	if err = h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	commentId, err := h.service.Create(userId, input)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"commentId": commentId,
	})
}

func (h *commentHandler) deleteComment(c *gin.Context) {
	// todo: when feedback is deleted -> delete related comments
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteComment not implemented",
	})
}

func (h *commentHandler) getCommentById(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "commentId is invalid",
		})
		return
	}

	if err = h.v.ValidateVar(commentId, "required,gt=0"); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "commentId is invalid",
		})
		return
	}

	comment, err := h.service.GetById(commentId)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "Comment not found",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *commentHandler) updateComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateComment not implemented",
	})
}
