package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/methods/auth"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/body"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if veer := validation.BodyValidation(c, &newUser); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	id := fmt.Sprintf("%d", len(users)+1)
	hashedPassword, errPass := auth.HashPassword(newUser.Password)
	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
			"error":   "Passsword to long",
		})
		return
	}

	newUser.Password = hashedPassword
	newUser.ID = id

	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
