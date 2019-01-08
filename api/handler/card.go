package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/internal/middleware"
)

// Card represents the Card handler
type Card struct {
	middleware middleware.Card
}

// NewCardHandler returns a newly created card handler
func NewCardHandler(ctx context.Context) (Handler, error) {
	middleware, err := middleware.NewCardMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	return &Card{middleware}, nil
}

// Create is the HTTP handler of the POST /card
func (m *Card) Create(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {

		userID := c.Param("userID")

		merchant, err := m.middleware.Create(ctx, userID)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "create card failed",
			})
			return
		}

		c.JSON(http.StatusCreated, merchant)
	}
}
