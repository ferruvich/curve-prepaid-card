package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Handler contains functions that should be
// Created inside each handler
type Handler interface {
	Create(context.Context) func(*gin.Context)
}

// ErrorMessage represents the generic error format
type ErrorMessage struct {
	Error string `json:"err"`
}
