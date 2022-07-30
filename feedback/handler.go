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
	Title    string  `json:"title" validate:"required,min=5,max=50"`
	Body     string  `json:"body" validate:"required,min=10,max=1000"`
	Category string  `json:"category" validate:"required,oneof=ui ux enchancement bug feature"`
	Status   *string `json:"status" validate:"omitempty,oneof=idea defined in-progress done"`
}

// swagger:route POST /api/feedback feedback createFeedback
// Create product feedback
// responses:
//	200: createFeedbackResponse

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

	if err := h.v.ValidateStruct(input); err != nil {
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

// swagger:route DELETE /api/feedback/:id feedback deleteFeedback
// Delete product feedback
// responses:
//	200: okResponse
//	404: notFoundResponse

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
		h.l.Error(err)
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

// swagger:route GET /api/feedback feedback getAllFeedback
// Returns all product feedback in the system
// responses:
//	200: getAllFeedbackResponse

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

// swagger:route GET /api/feedback/:id feedback getFeedbackById
// Returns product feedback by id
// responses:
//	200: getFeedbackByIdResponse
//	404: notFoundResponse

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

type updateFeedbackInput struct {
	Title    *string `json:"title" validate:"omitempty,min=5,max=50"`
	Body     *string `json:"body" validate:"omitempty,min=10,max=1000"`
	Category *string `json:"category" validate:"omitempty,oneof=ui ux enchancement bug feature"`
	Status   *string `json:"status" validate:"omitempty,oneof=idea defined in-progress done"`
}

// swagger:route PUT /api/feedback/:id feedback updateFeedback
// Returns product feedback by id
// responses:
//	200: okResponse
//	400: badRequestResponse
//	404: notFoundResponse

func (h *feedbackHandler) updateFeedback(c *gin.Context) {
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

	var input updateFeedbackInput
	if err := c.BindJSON(&input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid input",
		})
		return
	}

	if err = h.v.ValidateStruct(input); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid input",
		})
		return
	}

	err = h.service.Update(userId, feedbackIdInt, input)
	switch err {
	case nil:
		break
	case errNoInputToUpdate:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "No input to update",
		})
		return
	case sql.ErrNoRows:
		h.l.Error(err)
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
