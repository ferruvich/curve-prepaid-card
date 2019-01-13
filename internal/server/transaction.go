package server

import (
	"fmt"
	"net/http"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Transaction is the Transaction handler
type Transaction interface {
	GetTransactionList() func(*gin.Context)
}

// TransactionHandler is the Transaction struct
type TransactionHandler struct {
	server Server
}

// GetTransactionList return the transaction list of a given card
func (t *TransactionHandler) GetTransactionList() func(*gin.Context) {
	return func(c *gin.Context) {

		cardID := c.Param("cardID")

		txs, err := middleware.NewMiddleware(t.server.DataBase()).Transaction().GetListByCard(cardID)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusOK, txs)
	}
}
