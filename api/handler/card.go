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

// TopUpRequest embeds a topup request body
type TopUpRequest struct {
	Amount float64 `json:"amount" binding:"required"`
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

// TopUp is the HTTP handler of the POST /card/:id/topup
func (m *Card) TopUp(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {

		cardID := c.Param("cardID")

		request := &TopUpRequest{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "topup card failed",
			})
			return
		}

		err = m.middleware.TopUp(ctx, cardID, request.Amount)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "topup card failed",
			})
			return
		}

		c.Writer.WriteHeader(http.StatusNoContent)
	}
}
