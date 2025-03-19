package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/methods/task"
)

/*
  Setup all task routes
*/

func SetupTaskRoutes(routes *gin.Engine) {
	taskRoutes := routes.Group("/tasks")
	{
		taskRoutes.GET("/", task.GetAllTask)
		taskRoutes.GET("/:id", task.GetTaskByID)
		taskRoutes.POST("/", task.CreateTask)
		taskRoutes.PATCH("/:id", task.UpdateTask)
		taskRoutes.DELETE("/:id", task.DeleteTaskByID)
	}
}
