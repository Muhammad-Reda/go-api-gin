package main

import (
	"net/http"

	"github.com/muhammad-reda/go-api-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	routes.SetupUserRoutes(router)

	router.Run(":8080")
}
