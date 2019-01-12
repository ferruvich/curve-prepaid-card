package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// AuthorizationRequest represents the Authorization request handler
type AuthorizationRequest interface {
	Create() func(c *gin.Context)
}

// AuthorizationRequestHandler is the AuthorizationRequest struct
type AuthorizationRequestHandler struct {
	server Server
}

// AuthorizationRequestBody embeds a authorization request body
type AuthorizationRequestBody struct {
	MerchantID string  `json:"merchant_id" binding:"required"`
	CardID     string  `json:"card_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

// CaptureAuthorizationRequestBody embeds the
type CaptureAuthorizationRequestBody struct {
	Amount float64 `json:"amount" binding:"required"`
}

// Create is the HTTP handler of the POST /authorization
func (ar *AuthorizationRequestHandler) Create() func(c *gin.Context) {
	return func(c *gin.Context) {

		request := &AuthorizationRequestBody{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		card, err := middleware.NewMiddleware(ar.server.DataBase()).AuthorizationRequest().Create(
			request.MerchantID, request.CardID, request.Amount,
		)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusCreated, card)
	}
}
