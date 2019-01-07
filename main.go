package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/api/handler"
	"github.com/ferruvich/curve-challenge/internal/configuration"
)

func main() {
	cfg, err := configuration.GetConfiguration()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Building context
	ctx := context.WithValue(context.Background(), "cfg", cfg)

	// User routes
	userHandler, err := handler.NewUserHandler(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	router.POST("/user", userHandler.Create(ctx))

	// Merchant routes
	merchantHandler, err := handler.NewMerchantHandler(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	router.POST("/merchant", merchantHandler.Create(ctx))

	router.Run(strings.Join(
		[]string{
			cfg.Server.Host,
			cfg.Server.Port,
		}, ":",
	))
}
