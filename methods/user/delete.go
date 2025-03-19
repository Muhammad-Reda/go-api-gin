package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}
