package users

import "github.com/gin-gonic/gin"

func (h *UserHandler) AddRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/:id", h.getUser)
	}
}
