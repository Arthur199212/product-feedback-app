package feedback

import (
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FeedbackHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type feedbackHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service FeedbackService
}

func NewFeedbackHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service FeedbackService,
) FeedbackHandler {
	return &feedbackHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

func (h *feedbackHandler) getAllFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getAllFeedback not implemented",
	})
}

type createFeedbackInput struct {
	Title    string `json:"title" validate:"required,min=5,max=50"`
	Body     string `json:"body" validate:"required,min=10,max=1000"`
	Category string `json:"category" validate:"required,oneof=ui ux enchancement bug feature"`
	Status   string `json:"status" validate:"omitempty,oneof=idea defined in-progress done"`
}

func (h *feedbackHandler) createFeedback(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	var input createFeedbackInput
	if err := c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Input is invalid",
		})
		return
	}

	if err := h.v.Validate(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	feedbackId, err := h.service.Create(userId, input)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal service error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"feedbackId": feedbackId,
	})
}

func (h *feedbackHandler) getFeedbackById(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getFeedbackById not implemented",
	})
}

func (h *feedbackHandler) updateFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateFeedback not implemented",
	})
}

func (h *feedbackHandler) deleteFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "deleteFeedback not implemented",
	})
}
