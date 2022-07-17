package auth

import (
	users "product-feedback/user"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	userRepo := users.NewUserRepository()
	service := NewAuthService(userRepo)
	handler := NewAuthHandler(service)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handler.signIn)
		auth.POST("/sign-up", handler.signUp)
	}
}
