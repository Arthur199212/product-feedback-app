package users

import "github.com/gin-gonic/gin"

func (h *userHandler) AddRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/:id", h.getUser)
	}
}
