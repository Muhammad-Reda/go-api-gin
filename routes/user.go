package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/controller"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	"github.com/muhammad-reda/go-api-gin/domain/service"
)

var ctx *gin.Context

func UserApi(r *gin.Engine, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, ctx)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAll)
		v1.GET("/users/:id", userController.GetByID)
		v1.POST("/users", userController.Create)
		v1.PATCH("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)
	}
}
