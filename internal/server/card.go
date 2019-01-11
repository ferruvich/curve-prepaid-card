package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// Card represents the Card handler
type Card interface {
	Create() func(c *gin.Context)
	GetCard() func(c *gin.Context)
	Deposit() func(c *gin.Context)
}

// CardHandler is the Card struct
type CardHandler struct {
	server Server
}

// DepositRequest embeds a deposit request body
type DepositRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}

// Create is the HTTP handler of the POST /card
func (ch *CardHandler) Create() func(c *gin.Context) {
	return func(c *gin.Context) {

		userID := c.Param("userID")

		card, err := middleware.NewMiddleware(ch.server.DataBase()).Card().Create(userID)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "create card failed",
			})
			return
		}

		c.JSON(http.StatusCreated, card)
	}
}

// GetCard is the HTTP handler of the GET /card/:id
func (ch *CardHandler) GetCard() func(c *gin.Context) {
	return func(c *gin.Context) {

		cardID := c.Param("cardID")

		card, err := middleware.NewMiddleware(ch.server.DataBase()).Card().GetCard(cardID)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "get card failed",
			})
			return
		}

		c.JSON(http.StatusOK, card)
	}
}

// Deposit is the HTTP handler of the POST /card/:id/deposit
func (ch *CardHandler) Deposit() func(c *gin.Context) {
	return func(c *gin.Context) {

		cardID := c.Param("cardID")

		request := &DepositRequest{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: "bad request",
			})
			return
		}

		err = middleware.NewMiddleware(ch.server.DataBase()).Card().Deposit(cardID, request.Amount)
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
