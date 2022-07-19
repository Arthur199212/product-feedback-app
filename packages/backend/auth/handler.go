package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	AddRoutes(r *gin.RouterGroup)
}

type authHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) AuthHandler {
	return &authHandler{service}
}

func (h *authHandler) signIn(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "signIn not implemented",
	})
}

func (h *authHandler) signUp(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotImplemented, map[string]interface{}{
		"message": "signUp not implemented",
	})
}
