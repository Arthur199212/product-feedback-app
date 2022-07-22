package user

import (
	"database/sql"
	"net/http"
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{service}
}

func (h *userHandler) getUser(c *gin.Context) {
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
			"message": "Internal service error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
