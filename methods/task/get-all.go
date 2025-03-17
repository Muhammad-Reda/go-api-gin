package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/dummy"
)

var tasks = dummy.Tasks

func GetAllTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    tasks,
	})
}
