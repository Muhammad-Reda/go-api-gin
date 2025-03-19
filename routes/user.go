package routes

import (
	"github.com/muhammad-reda/go-api-gin/methods/user"

	"github.com/gin-gonic/gin"
)

/*
	Setup all user routes
*/

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", user.GetAllUsers)
		userRoutes.GET("/:id", user.GetUserByID)
		userRoutes.POST("/", user.CreateUser)
		userRoutes.POST("/login", user.Login)
		userRoutes.PATCH("/:id", user.UpdateUserByID)
		userRoutes.DELETE("/:id", user.DeleteUserByID)
	}
}
