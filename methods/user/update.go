package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
)

func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.UpdateUser

	if veer := validation.BodyValidation(c, &updatedUser); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id

			if updatedUser.Email == "" {
				updatedUser.Email = user.Email
			}
			if updatedUser.Password == "" {
				updatedUser.Password = user.Password
			}
			if updatedUser.Username == "" {
				updatedUser.Username = user.Username
			}

			users[i] = models.User{
				ID:       updatedUser.ID,
				Username: updatedUser.Username,
				Password: updatedUser.Password,
				Email:    updatedUser.Email,
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "User updated",
			})

			return

		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "User not found",
	})
}
