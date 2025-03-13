package routes

import (
	"github.com/muhammad-reda/go-api-gin/methods"

	"github.com/gin-gonic/gin"
)

/*
	Setup all user routes
*/

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", methods.GetAllUsers)
		userRoutes.GET("/:id", methods.GetUserById)
		userRoutes.POST("/", methods.CreateUser)
		userRoutes.PUT("/:id", methods.UpdateUserById)
		userRoutes.DELETE("/:id", methods.DeleteUserById)
	}
}
