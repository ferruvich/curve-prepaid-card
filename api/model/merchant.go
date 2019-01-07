package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Merchant represents a Merchant inside system
type Merchant struct {
	ID string `json:"ID"`
}

// NewMerchant returns a newly created merchant
func NewMerchant() (*Merchant, error) {

	merchantUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating merchant id")
	}

	return &Merchant{
		ID: merchantUUID.String(),
	}, nil
}
