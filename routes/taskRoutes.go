package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/domain/repository"
	"github.com/muhammad-reda/go-api-gin/domain/service"
	"github.com/muhammad-reda/go-api-gin/handler"
)

func TaskApi(routes *gin.Engine, DB *sql.DB) {
	userRepo := repository.NewUserRepository(DB)
	taskRepository := repository.NewTaskRepository(DB)
	taskService := service.NewTaskService(taskRepository, userRepo)
	taskHandler := handler.NewTaskHandler(taskService, ctx)

	v1 := routes.Group("/api/v1")
	{
		v1.GET("/tasks", taskHandler.GetAllTask)
		v1.GET("/tasks/:id", taskHandler.GetTaskById)
		v1.POST("/tasks", taskHandler.Create)
		v1.PATCH("/tasks/:id", taskHandler.Update)
		v1.DELETE("/tasks/:id", taskHandler.Delete)
	}
}
