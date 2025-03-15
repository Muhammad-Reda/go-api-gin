package methods

import (
	"fmt"
	"net/http"

	"github.com/muhammad-reda/go-api-gin/dummy"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/user"

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
	var updatedUser models.UpdateUser

	if veer := validation.UserValidation(c, nil, &updatedUser); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			if updatedUser.Username == "" {
				updatedUser.Username = user.Username
			}
			if updatedUser.Email == "" {
				updatedUser.Email = user.Email
			}
			if updatedUser.Password == "" {
				updatedUser.Password = user.Password
			}
			if updatedUser.Age == 0 {
				updatedUser.Age = user.Age
			}
			if updatedUser.Address == "" {
				updatedUser.Address = user.Address
			}
			if updatedUser.Phone == "" {
				updatedUser.Phone = user.Phone
			}

			users[i] = models.User{
				ID:       id,
				Username: updatedUser.Username,
				Password: updatedUser.Password,
				Email:    updatedUser.Email,
				Age:      updatedUser.Age,
				Address:  updatedUser.Address,
				Phone:    updatedUser.Phone,
			}

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

	// Validation request body
	if verr := validation.UserValidation(c, &newUser, nil); len(verr) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body",
			"errors":  verr,
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
