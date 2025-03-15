package main

import (
	"net/http"

	"github.com/muhammad-reda/go-api-gin/routes"
	"github.com/muhammad-reda/go-api-gin/validation/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.CorsMiddleware())
	// router.SetTrustedProxies([]string{"7.0.0.2"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
			"origin":  c.Request.Header.Get("Origin"),
		})
	})

	routes.SetupUserRoutes(router)

	router.Run(":8080")
}
