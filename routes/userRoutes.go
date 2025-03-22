package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	"github.com/muhammad-reda/go-api-gin/domain/service"
	"github.com/muhammad-reda/go-api-gin/handler"
)

var ctx *gin.Context

func UserApi(r *gin.Engine, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, ctx)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetAll)
		v1.GET("/users/:id", userHandler.GetByID)
		v1.POST("/users", userHandler.Create)
		v1.PATCH("/users/:id", userHandler.Update)
		v1.DELETE("/users/:id", userHandler.Delete)
	}
}
