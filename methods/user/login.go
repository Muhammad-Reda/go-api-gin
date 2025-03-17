package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/methods/auth"
	"github.com/muhammad-reda/go-api-gin/models"
	validation "github.com/muhammad-reda/go-api-gin/validation/user"
)

func Login(c *gin.Context) {
	var userLogin models.Loginuser

	if veer := validation.UserValidation(c, &userLogin); len(veer) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": veer,
		})
		return
	}

	for i, user := range users {
		if user.Username == userLogin.Username {
			matchedPass, err := auth.CheckPassword(userLogin.Password, users[i].Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Username atau password salah",
				})
				return
			}

			if matchedPass {
				c.JSON(http.StatusOK, gin.H{
					"message": "Login success",
				})
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"messsage": "User not found",
	})
}
