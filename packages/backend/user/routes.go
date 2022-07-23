package user

import (
	"product-feedback/middleware"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) AddRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users", middleware.AuthRequired)
	{
		users.GET("/me", h.getMe)
		users.GET("/:id", h.getUserById)
	}
}
