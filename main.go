package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/handler"
)

func main() {
	cfg := configuration.GetConfiguration()

	if cfg.Server.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Building context
	ctx := context.WithValue(context.Background(), "cfg", cfg)

	// Initializing handlers

	// User handler
	userHandler, err := handler.NewUserHandler(ctx)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Merchant Handler
	merchantHandler, err := handler.NewMerchantHandler(ctx)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// Card Handler
	cardHandler, err := handler.NewCardHandler(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	// AuthorizationRequest handler
	authReqHandler, err := handler.NewAuthoziationRequestHandler(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	// Routes
	router.POST("/user", userHandler.Create(ctx))
	router.POST("/user/:userID/card", cardHandler.Create(ctx))
	router.GET("/user/:userID/card/:cardID", cardHandler.GetCard(ctx))
	router.POST("/user/:userID/card/:cardID/deposit", cardHandler.Deposit(ctx))
	router.POST("/merchant", merchantHandler.Create(ctx))
	router.POST("/authorization", authReqHandler.Create(ctx))

	router.Run(strings.Join(
		[]string{
			cfg.Server.Host,
			cfg.Server.Port,
		}, ":",
	))
}
