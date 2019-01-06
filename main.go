package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/api/handler"
)

func main() {
	router := gin.Default()

	// User routes
	userHandler, err := handler.NewUserHandler()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	router.POST("/user", userHandler.Create)

	router.Run(":8080")
}
