package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    task,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task not found",
	})
}
