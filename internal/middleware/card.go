package middleware

import (
	"context"

	"github.com/ferruvich/curve-challenge/internal/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
)

// Card represents the card middleware interface
type Card interface {
	Create(context.Context, string) (*model.Card, error)
	SetUserMiddleware(User)
}

// CardMiddleware is the Card implementation
type CardMiddleware struct {
	repo           repo.Card
	userMiddleware User
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

// SetUserMiddleware sets the user middleware in order to use it to
// make operations on user
func (m *CardMiddleware) SetUserMiddleware(userMiddleware User) {
	m.userMiddleware = userMiddleware
}

// Create creates and returns new merchants, or an error if there is one
func (m *CardMiddleware) Create(ctx context.Context, ownerID string) (*model.Card, error) {

	user, err := m.userMiddleware.Read(ctx, ownerID)
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
