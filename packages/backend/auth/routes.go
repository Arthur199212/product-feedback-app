package auth

import (
	"github.com/gin-gonic/gin"
)

func (h *authHandler) AddRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.GET("/github", h.redirectToGitHubLoginURL)
		auth.GET("/github/callback", h.loginWithGitHub)
		auth.POST("/refresh-token", h.refreshAccessToken)
	}
}
