package methods

import (
	"fmt"
	"net/http"

	"github.com/muhammad-reda/go-api-gin/dummy"
	"github.com/muhammad-reda/go-api-gin/models"

	"github.com/gin-gonic/gin"
)

var users = dummy.Users

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if errBodyJson := c.ShouldBindBodyWithJSON(&updatedUser); errBodyJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated",
				"data":    updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if errBodyJson := c.ShouldBindBodyWithJSON(&newUser); errBodyJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
		})
		return
	}

	// Generate new ID
	newUser.ID = fmt.Sprintf("%d", len(users)+1)

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"data":    newUser,
	})

}

func DeleteUserById(c *gin.Context) {
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
