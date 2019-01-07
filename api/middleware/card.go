package middleware

import (
	"context"

	"github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
)

// Card represents the card middleware interface
type Card interface {
	Create(context.Context, string, User) (*model.Card, error)
}

// CardMiddleware is the Card implementation
type CardMiddleware struct {
	repo repo.Card
}

// NewCardMiddleware returns a new middleware for user
func NewCardMiddleware(ctx context.Context) (Card, error) {
	repo, err := repo.NewCardRepo(ctx)
	if err != nil {
		return nil, err
	}

	return &CardMiddleware{
		repo: repo,
	}, nil
}

// Create creates and returns new merchants, or an error if there is one
func (m *CardMiddleware) Create(ctx context.Context, ownerID string, userMiddleware User) (*model.Card, error) {

	user, err := userMiddleware.Read(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	card, err := model.NewCard(user.ID)
	if err != nil {
		return nil, err
	}

	if err = m.repo.Write(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}
