package main

import (
	"log"
	"os"

	"github.com/gasuhwbab/goAuth/internal/controller/user"
	"github.com/gasuhwbab/goAuth/internal/initializer"
	"github.com/gasuhwbab/goAuth/internal/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.InitDb()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	router.POST("/signup", user.Signup)
	router.POST("/login", user.Login)
	router.GET("/validate", middleware.RequireAuth, user.Validate)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
