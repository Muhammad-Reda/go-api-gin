package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
)

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var taskUpdate models.UpdateTask

	if veer := validation.BodyValidation(c, &taskUpdate); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			taskUpdate.ID = id

			if taskUpdate.Name == "" {
				taskUpdate.Name = task.Name
			}

			if taskUpdate.Status == "" {
				taskUpdate.Status = task.Status
			}

			tasks[i] = models.Task{
				ID:     taskUpdate.ID,
				Name:   taskUpdate.Name,
				Status: taskUpdate.Status,
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Task updated",
				"data":    taskUpdate,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task not found",
	})
}
