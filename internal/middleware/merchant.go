package middleware

import (
	"context"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

// Merchant represents the merchant middleware interface
type Merchant interface {
	Create(context.Context) (*model.Merchant, error)
}

// MerchantMiddleware is the Merchant implementation
type MerchantMiddleware struct {
	database database.Merchant
}

// NewMerchantMiddleware returns a new middleware for user
func NewMerchantMiddleware(ctx context.Context) (Merchant, error) {
	database, err := database.NewMerchantdatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &MerchantMiddleware{
		database: database,
	}, nil
}

// Create creates and returns new merchants, or an error if there is one
func (m *MerchantMiddleware) Create(ctx context.Context) (*model.Merchant, error) {

	merchant, err := model.NewMerchant()
	if err != nil {
		return nil, err
	}

	if err = m.database.Write(ctx, merchant); err != nil {
		return nil, err
	}

	return merchant, nil
}
