package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/user"
)

func UpdateUserByID(c *gin.Context) {
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

			if updatedUser.Address == "" {
				updatedUser.Address = user.Address
			}
			if updatedUser.Email == "" {
				updatedUser.Email = user.Email
			}
			if updatedUser.Password == "" {
				updatedUser.Password = user.Password
			}
			if updatedUser.Phone == "" {
				updatedUser.Phone = user.Phone
			}
			if updatedUser.Username == "" {
				updatedUser.Username = user.Username
			}
			if updatedUser.Age == 0 {
				updatedUser.Age = user.Age
			}

			users[i] = models.User{
				ID:       updatedUser.ID,
				Username: updatedUser.Username,
				Password: updatedUser.Password,
				Email:    updatedUser.Email,
				Age:      updatedUser.Age,
				Address:  updatedUser.Address,
				Phone:    updatedUser.Phone,
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
