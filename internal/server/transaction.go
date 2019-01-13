package server

import (
	"fmt"
	"net/http"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Transaction is the Transaction handler
type Transaction interface {
	Create() func(*gin.Context)
}

// TransactionHandler is the Transaction struct
type TransactionHandler struct {
	server Server
}

// CaptureBody binds the /capture body
type CaptureBody struct {
	Amount float64 `json:"amount" binding:"required"`
}

// Create creates a new transaction
func (t *TransactionHandler) Create() func(*gin.Context) {
	return func(c *gin.Context) {

		authReqID := c.Param("authID")

		request := &CaptureBody{}
		err := c.BindJSON(request)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusBadRequest, ErrorMessage{
				Error: "bad request",
			})
			return
		}

		tx, err := middleware.NewMiddleware(t.server.DataBase()).Transaction().CreatePayment(authReqID, request.Amount)
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
