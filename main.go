package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/api/handler"
)

func main() {
	router := gin.Default()

	// User routes
	userHandler := handler.NewUserHandler()
	router.POST("/user", userHandler.Create)

	router.Run(":8080")
}
