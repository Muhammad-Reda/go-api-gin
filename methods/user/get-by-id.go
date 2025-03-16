package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"data":    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User Not found",
	})
}
