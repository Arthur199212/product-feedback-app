package user

import (
	"database/sql"
	"net/http"
	"product-feedback/middleware"
	"product-feedback/validation"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type userHandler struct {
	l       *logrus.Logger
	v       *validation.Validation
	service UserService
}

func NewUserHandler(
	l *logrus.Logger,
	v *validation.Validation,
	service UserService,
) UserHandler {
	return &userHandler{
		l:       l,
		v:       v,
		service: service,
	}
}

// swagger:route GET /api/users/me users getMe
// Returs user data
//
// security:
// - Bearer:
//
// responses:
//	200: getUserResponse

func (h *userHandler) getMe(c *gin.Context) {
	userId, err := middleware.GetUserIdFromGinCtx(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
	}

	user, err := h.service.GetById(userId)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// swagger:route GET /api/users/:id users getUserById
// Returs user data by id
//
// security:
// - Bearer:
//
// responses:
//	200: getUserResponse
//	404: errorResponse

func (h *userHandler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userId is invalid",
		})
		return
	}

	if err = h.v.ValidateVar(userId, "required,gt=0"); err != nil {
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "userId is invalid",
		})
		return
	}

	user, err := h.service.GetById(userId)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
		return
	default:
		h.l.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
