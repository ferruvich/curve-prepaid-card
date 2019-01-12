package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// AuthorizationRequest represents the request to authorize a payment
type AuthorizationRequest struct {
	ID       string  `json:"ID"`
	Merchant string  `json:"merchant"`
	Card     string  `json:"card"`
	Approved bool    `json:"approved"`
	Amount   float64 `json:"amount"`
	Reversed float64 `json:"reversed"`
}

// NewAuthorizationRequest returns a newly created AuthorizationRequest
func NewAuthorizationRequest(merchantID string, cardID string, amount float64) (*AuthorizationRequest, error) {
	authID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating authorization id")
	}

	return &AuthorizationRequest{
		ID: authID.String(), Merchant: merchantID, Card: cardID,
		Approved: false, Amount: amount, Reversed: 0.0,
	}, nil
}

// Revert reverse some of the amount of the authorization request
func (ar *AuthorizationRequest) Revert(amount float64) error {
	if ar.Reversed+amount > ar.Amount {
		return errors.Errorf("cannot reverse such amount")
	}

	ar.Reversed += amount

	return nil
}

// Approve is called when an authorization is approved
func (ar *AuthorizationRequest) Approve() {
	ar.Approved = true
}
