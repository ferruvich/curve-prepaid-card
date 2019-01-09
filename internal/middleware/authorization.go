package middleware

import (
	"context"

	"github.com/ferruvich/curve-prepaid-card/internal/repo"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

// AuthorizationRequest represents the middleware of authorization requests
type AuthorizationRequest interface {
	Create(context.Context, string, string, float64) (*model.AuthorizationRequest, error)
}

// AuthorizationRequestMiddleware is the AuthorizationRequest implementation
type AuthorizationRequestMiddleware struct {
	repo repo.AuthorizationRequest
}

// NewAuthorizationRequestMiddleware returns a new AuthorizationRequest
func NewAuthorizationRequestMiddleware(ctx context.Context) (AuthorizationRequest, error) {
	repo, err := repo.NewAuthorizationRequestRepo(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthorizationRequestMiddleware{
		repo: repo,
	}, nil
}

// Create creates and returns a new authorization request
func (ar *AuthorizationRequestMiddleware) Create(
	ctx context.Context, merchantID string, cardID string, amount float64,
) (*model.AuthorizationRequest, error) {

	authReq, err := model.NewAuthorizationRequest(
		merchantID, cardID, amount,
	)
	if err != nil {
		return nil, err
	}

	if err = ar.repo.Write(ctx, authReq); err != nil {
		return nil, err
	}

	return authReq, nil

}
