package main

import (
	"log"
	"net/http"
	"os"

	"github.com/muhammad-reda/go-api-gin/config"
	"github.com/muhammad-reda/go-api-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	db, _ := config.GetDB()

	// Load env
	errEnv := godotenv.Load()
	appPort := os.Getenv("APP_PORT")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// router.Use(cors.CorsMiddleware())
	// router.SetTrustedProxies([]string{"7.0.0.2"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
			"db":      db,
		})
	})

	routes.UserApi(router, db)

	router.Run(":" + appPort)
}
