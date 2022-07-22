package feedback

import (
	"database/sql"
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"
	"strconv"

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
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"feedbackId": feedbackId,
	})
}

func (h *feedbackHandler) deleteFeedback(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	feedbackIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid feedback id",
		})
		return
	}

	err = h.service.Delete(userId, feedbackIdInt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "Feedback not found",
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

func (h *feedbackHandler) getAllFeedback(c *gin.Context) {
	// todo: implement options:
	// filter by: userId, category, status
	// filter by group: status=idea&status=default&category=ui&category=ux
	// sorted: most voted, least voted, most commented, least commented
	// pagination: limit=<uint>, page=<uint>
	fList, err := h.service.GetAll()
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, fList)
}

func (h *feedbackHandler) getFeedbackById(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	feedbackIdInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid feedback id",
		})
		return
	}

	feedback, err := h.service.GetById(userId, feedbackIdInt)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "Feedback not found",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func (h *feedbackHandler) updateFeedback(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "updateFeedback not implemented",
	})
}
