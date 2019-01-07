package middleware

import (
	"context"

	"github.com/ferruvich/curve-challenge/internal/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
)

// Merchant represents the merchant middleware interface
type Merchant interface {
	Create(context.Context) (*model.Merchant, error)
}

// MerchantMiddleware is the Merchant implementation
type MerchantMiddleware struct {
	repo repo.Merchant
}

// NewMerchantMiddleware returns a new middleware for user
func NewMerchantMiddleware(ctx context.Context) (Merchant, error) {
	repo, err := repo.NewMerchantRepo(ctx)
	if err != nil {
		return nil, err
	}

	return &MerchantMiddleware{
		repo: repo,
	}, nil
}

// Create creates and returns new merchants, or an error if there is one
func (m *MerchantMiddleware) Create(ctx context.Context) (*model.Merchant, error) {

	merchant, err := model.NewMerchant()
	if err != nil {
		return nil, err
	}

	if err = m.repo.Write(ctx, merchant); err != nil {
		return nil, err
	}

	return merchant, nil
}
