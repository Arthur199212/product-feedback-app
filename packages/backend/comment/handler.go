package comment

import (
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"

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

	// todo: filter by feedback
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllComments not implemented",
	})
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

func (h *commentHandler) updateComment(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateComment not implemented",
	})
}

func (h *commentHandler) deleteComment(c *gin.Context) {
	// todo: when feedback is deleted -> delete related comments
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteComment not implemented",
	})
}
