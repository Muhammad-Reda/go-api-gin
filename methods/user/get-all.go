package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/muhammad-reda/go-api-gin/dummy"
)

var users = dummy.Users

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})
}
