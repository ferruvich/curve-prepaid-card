package middleware

import (
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=authorization_mock.go -source=authorization.go -package=middleware -self_package=. AuthorizationRequest

// AuthorizationRequest represents the middleware of authorization requests
type AuthorizationRequest interface {
	Create(string, string, float64) (*model.AuthorizationRequest, error)
}

// AuthorizationRequestMiddleware is the AuthorizationRequest implementation
type AuthorizationRequestMiddleware struct {
	middleware Middleware
}

// Create creates and returns a new authorization request
func (ar *AuthorizationRequestMiddleware) Create(merchantID string, cardID string, amount float64) (*model.AuthorizationRequest, error) {

	// Amount blocked, creating authorization
	authReq, err := model.NewAuthorizationRequest(
		merchantID, cardID, amount,
	)
	if err != nil {
		return nil, err
	}

	// Getting card to block amount
	card, err := ar.middleware.DataBase().Card().Read(cardID)
	if err != nil {
		return nil, err
	}

	if err = card.BlockAmount(amount); err != nil {
		return nil, err
	}

	// Since blockAmount returns no error, we approve the auth request
	authReq.Approve()

	// Updating card and writing authorization
	if err = ar.middleware.DataBase().Card().Update(card); err != nil {
		return nil, err
	}

	if err = ar.middleware.DataBase().AuthorizationRequest().Write(authReq); err != nil {
		return nil, err
	}

	return authReq, nil

}
