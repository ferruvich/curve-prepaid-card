package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// AuthorizationRequest represents the Card handler
type AuthorizationRequest struct {
	middleware middleware.AuthorizationRequest
}

// AuthorizationRequestBody embeds a topup request body
type AuthorizationRequestBody struct {
	MerchantID string  `json:"merchant_id" binding:"required"`
	CardID     string  `json:"card_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

// NewAuthoziationRequestHandler returns a newly created authorization request handler
func NewAuthoziationRequestHandler(ctx context.Context) (*AuthorizationRequest, error) {
	middleware, err := middleware.NewAuthorizationRequestMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthorizationRequest{middleware}, nil
}

// Create is the HTTP handler of the POST /card
func (m *AuthorizationRequest) Create(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {

		request := &AuthorizationRequestBody{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: "bad request",
			})
			return
		}

		card, err := m.middleware.Create(ctx, request.MerchantID, request.CardID, request.Amount)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "create authorization request failed",
			})
			return
		}

		c.JSON(http.StatusCreated, card)
	}
}
