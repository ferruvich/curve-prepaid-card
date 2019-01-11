package middleware

import (
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

// Card represents the card middleware interface
type Card interface {
	Create(string) (*model.Card, error)
	GetCard(string) (*model.Card, error)
	Deposit(string, float64) error
}

// CardMiddleware is the Card implementation
type CardMiddleware struct {
	middleware Middleware
}

// Create creates and returns new merchants, or an error if there is any
func (c *CardMiddleware) Create(ownerID string) (*model.Card, error) {

	card, err := model.NewCard(ownerID)
	if err != nil {
		return nil, err
	}

	if err = c.middleware.DataBase().Card().Write(card); err != nil {
		return nil, err
	}

	return card, nil
}

// GetCard gets and returns an existing card, or an error if there is any
func (c *CardMiddleware) GetCard(cardID string) (*model.Card, error) {

	card, err := c.middleware.DataBase().Card().Read(cardID)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// Deposit topups a card, adding some money
func (c *CardMiddleware) Deposit(cardID string, amount float64) error {

	card, err := c.middleware.DataBase().Card().Read(cardID)
	if err != nil {
		return err
	}

	card.IncrementAccountBalance(amount)

	if err = c.middleware.DataBase().Card().Update(card); err != nil {
		return err
	}

	return nil
}
