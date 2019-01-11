package middleware

import (
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=merchant_mock.go -source=merchant.go -package=middleware -self_package=. Merchant

// Merchant represents the merchant middleware interface
type Merchant interface {
	Create() (*model.Merchant, error)
}

// MerchantMiddleware is the Merchant implementation
type MerchantMiddleware struct {
	middleware Middleware
}

// Create creates and returns new merchants, or an error if there is one
func (m *MerchantMiddleware) Create() (*model.Merchant, error) {

	merchant, err := model.NewMerchant()
	if err != nil {
		return nil, err
	}

	if err = m.middleware.DataBase().Merchant().Write(merchant); err != nil {
		return nil, err
	}

	return merchant, nil
}
