package routes

import (
	"net/http"

	"github.com/muhammad-reda/go-api-gin/db"
	"github.com/muhammad-reda/go-api-gin/methods/user"
	"github.com/muhammad-reda/go-api-gin/models"

	"github.com/gin-gonic/gin"
)

/*
	Setup all user routes
*/

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", user.GetAllUsers)
		userRoutes.GET("/:id", user.GetUserByID)
		userRoutes.POST("/", user.CreateUser)
		userRoutes.POST("/login", user.Login)
		userRoutes.PATCH("/:id", user.UpdateUserByID)
		userRoutes.DELETE("/:id", user.DeleteUserByID)
		userRoutes.GET("/db", func(c *gin.Context) {
			db, err := db.Connection()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
					"error":   err.Error(),
				})
				return
			}
			defer db.Close()

			var tasks []models.Task
			query := "SELECT * FROM tasks"

			rows, errDb := db.Query(query)
			if errDb != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
					"error":   err.Error(),
				})
			}

			for rows.Next() {
				var task models.Task
				errScan := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Status, &task.UserID, &task.CreatedAT, &task.UpdatedAt)
				if errScan != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Internal server error",
						"code":    errScan.Error(),
					})
					return
				}
				tasks = append(tasks, task)
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "db connected",
				"db":      tasks,
			})
		})
	}
}
