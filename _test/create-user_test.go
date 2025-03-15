package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammad-reda/go-api-gin/methods"
	"github.com/muhammad-reda/go-api-gin/models"
)

func setupRouter() *gin.Engine {
	// Set to test mode
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	router.POST("/users", methods.CreateUser)

	return router
}

func BenchmarkCreateUser(b *testing.B) {
	router := setupRouter()

	// dummy data for testing
	user := models.User{
		Username: "testUser",
		Password: "pass",
		Email:    "test@test.com",
		Age:      22,
		Address:  "Indonesia",
		Phone:    123213,
	}

	jsonValue, _ := json.Marshal(user)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		// Create response recorder
		w := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(w, req)

		// validate the response
		if w.Code != http.StatusCreated {
			b.Fatalf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	}
}

func BenchmarkGetUsers(b *testing.B) {
	router := setupRouter()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

	}
}
