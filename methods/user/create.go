package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/user"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if veer := validation.UserValidation(c, &newUser, nil); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	id := fmt.Sprintf("%d", len(users)+1)
	newUser.ID = id

	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
