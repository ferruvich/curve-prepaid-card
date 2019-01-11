package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// Merchant represents the Merchant handler
type Merchant interface {
	Create() func(c *gin.Context)
}

// MerchantHandler is the Merchant struct
type MerchantHandler struct {
	server Server
}

// Create is the HTTP handler of the POST /merchant
func (m *MerchantHandler) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		merchant, err := middleware.NewMiddleware(m.server.DataBase()).Merchant().Create()
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "create merchant failed",
			})
			return
		}

		c.JSON(http.StatusCreated, merchant)
	}
}
