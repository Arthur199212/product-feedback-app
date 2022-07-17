package users

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	users := rg.Group("/users")
	{
		users.GET("/:id", handler.getUser)
	}
}
