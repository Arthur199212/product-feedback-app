package users

import (
	"net/http"

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
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "getUser not implemented",
	})
}
