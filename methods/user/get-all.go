package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/muhammad-reda/go-api-gin/db"
	"github.com/muhammad-reda/go-api-gin/dummy"
	"github.com/muhammad-reda/go-api-gin/models"
)

var users = dummy.Users

func GetAllUsers(c *gin.Context) {
	db, errDb := db.Connection()
	if errDb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"code":    "Connection to database",
		})
		return
	}
	defer db.Close()

	var users []models.User
	query := "SELECT id, email, username, password, created_at, updated_at FROM users"

	// Take context from gin
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rows, errQuery := db.QueryContext(ctx, query)
	if errQuery != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"code":    "query database",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		errScan := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if errScan != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"code":    "Scan Query",
			})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})
}
