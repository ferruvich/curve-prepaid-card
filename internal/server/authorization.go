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
	Capture() func(c *gin.Context)
	Revert() func(c *gin.Context)
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

// AmountBody embeds a revert request, or a capture one
type AmountBody struct {
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

		authRequest, err := middleware.NewMiddleware(ar.server.DataBase()).AuthorizationRequest().Create(
			request.MerchantID, request.CardID, request.Amount,
		)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusCreated, authRequest)
	}
}

// Capture captures an auth request, is used in POST /authorization/:id/capture
func (ar *AuthorizationRequestHandler) Capture() func(*gin.Context) {
	return func(c *gin.Context) {

		authReqID := c.Param("authID")

		request := &AmountBody{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: "bad request",
			})
			return
		}

		tx, err := middleware.NewMiddleware(ar.server.DataBase()).Transaction().CreatePayment(authReqID, request.Amount)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusCreated, tx)

	}
}

// Revert reverts some amount of an auth request, is used in POST /authorization/:id/revert
func (ar *AuthorizationRequestHandler) Revert() func(c *gin.Context) {
	return func(c *gin.Context) {

		authID := c.Param("authID")

		request := &AmountBody{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		if err = middleware.NewMiddleware(ar.server.DataBase()).AuthorizationRequest().Revert(
			authID, request.Amount,
		); err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		c.Writer.WriteHeader(http.StatusAccepted)
	}
}
