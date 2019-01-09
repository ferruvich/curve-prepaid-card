package middleware

import (
	"context"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

// Card represents the card middleware interface
type Card interface {
	Create(context.Context, string) (*model.Card, error)
	GetCard(context.Context, string) (*model.Card, error)
	Deposit(context.Context, string, float64) error
}

// CardMiddleware is the Card implementation
type CardMiddleware struct {
	database database.Card
}

// NewCardMiddleware returns a new middleware for user
func NewCardMiddleware(ctx context.Context) (Card, error) {
	database, err := database.NewCarddatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &CardMiddleware{
		database: database,
	}, nil
}

// Create creates and returns new merchants, or an error if there is any
func (m *CardMiddleware) Create(ctx context.Context, ownerID string) (*model.Card, error) {

	card, err := model.NewCard(ownerID)
	if err != nil {
		return nil, err
	}

	if err = m.database.Write(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}

// GetCard gets and returns an existing card, or an error if there is any
func (m *CardMiddleware) GetCard(ctx context.Context, cardID string) (*model.Card, error) {

	card, err := m.database.Read(ctx, cardID)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// Deposit topups a card, adding some money
func (m *CardMiddleware) Deposit(ctx context.Context, cardID string, amount float64) error {

	card, err := m.database.Read(ctx, cardID)
	if err != nil {
		return err
	}

	card.IncrementAccountBalance(amount)

	if err = m.database.Update(ctx, card); err != nil {
		return err
	}

	return nil
}
