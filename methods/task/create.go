package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
)

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if veer := validation.BodyValidation(c, &newTask); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	id := fmt.Sprintf("%d", len(tasks)+1)
	newTask.ID = id

	tasks = append(tasks, newTask)

	c.JSON(http.StatusOK, gin.H{
		"message": "Task created",
	})
}
