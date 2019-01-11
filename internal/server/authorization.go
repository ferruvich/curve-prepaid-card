package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// AuthorizationRequest represents the Card handler
type AuthorizationRequest interface {
	Create() func(c *gin.Context)
}

// AuthorizationRequestHandler is the AuthorizationRequest struct
type AuthorizationRequestHandler struct {
	server Server
}

// AuthorizationRequestBody embeds a topup request body
type AuthorizationRequestBody struct {
	MerchantID string  `json:"merchant_id" binding:"required"`
	CardID     string  `json:"card_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

// Create is the HTTP handler of the POST /card
func (ar *AuthorizationRequestHandler) Create() func(c *gin.Context) {
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

		card, err := middleware.NewMiddleware(ar.server.DataBase()).AuthorizationRequest().Create(
			request.MerchantID, request.CardID, request.Amount,
		)
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
