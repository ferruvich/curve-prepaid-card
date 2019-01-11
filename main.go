package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/server"
)

func main() {
	cfg := configuration.GetConfiguration()

	if cfg.Server.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Routes
	router.POST("/user", server.NewUserHandler().Create())
	router.POST("/user/:userID/card", server.NewCardHandler().Create())
	router.GET("/user/:userID/card/:cardID", server.NewCardHandler().GetCard())
	router.POST("/user/:userID/card/:cardID/deposit", server.NewCardHandler().Deposit())
	router.POST("/merchant", server.NewMerchantHandler().Create())
	router.POST("/authorization", server.NewAuthorizationRequestHandler().Create())

	router.Run(strings.Join(
		[]string{
			cfg.Server.Host,
			cfg.Server.Port,
		}, ":",
	))
}
