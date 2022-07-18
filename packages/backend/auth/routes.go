package auth

import (
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) AddRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}
}
