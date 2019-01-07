package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/api/middleware"
)

// Merchant represents the User handler
type Merchant struct {
	middleware middleware.Merchant
}

// NewUserHandler returns a newly created user handler
func NewMerchantHandler(ctx context.Context) (Handler, error) {
	middleware, err := middleware.NewMerchantMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	return &Merchant{middleware}, nil
}

// Create is the HTTP handler of the POST /user
func (u *Merchant) Create(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		merchant, err := u.middleware.Create(ctx)
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
