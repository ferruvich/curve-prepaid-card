package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// Merchant represents the Merchant handler
type Merchant struct {
	middleware middleware.Merchant
}

// NewMerchantHandler returns a newly created merchant handler
func NewMerchantHandler(ctx context.Context) (*Merchant, error) {
	middleware, err := middleware.NewMerchantMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	return &Merchant{middleware}, nil
}

// Create is the HTTP handler of the POST /merchant
func (m *Merchant) Create(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		merchant, err := m.middleware.Create(ctx)
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
