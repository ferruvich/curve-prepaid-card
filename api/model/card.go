package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Card represents a new card
type Card struct {
	ID               string  `json:"ID"`
	Owner            string  `json:"owner"`
	AccountBalance   float64 `json:"account_balance"`
	AvailableBalance float64 `json:"available_balance"`
}

// NewCard returns a newly created card, owned by the user of given ID
func NewCard(userID string) (*Card, error) {
	cardUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating user id")
	}

	return &Card{
		ID:               cardUUID.String(),
		Owner:            userID,
		AccountBalance:   0.0,
		AvailableBalance: 0.0,
	}, nil
}

// IncrementAccountBalance increments the available balance
func (c *Card) IncrementAccountBalance(amount float64) {
	c.AccountBalance += amount

	/* Account balance is incremented due to topups or refunds
	 * since this amount is always spendable
	 * we are intrementing available balance too */
	c.AvailableBalance += amount
}

// ReverseAmountBlocked decrements the available balance
// thanks to a reverse in an authorization request
func (c *Card) ReverseAmountBlocked(amount float64) error {
	if c.AvailableBalance+amount > c.AccountBalance {
		return errors.Errorf("available balance cannot be greater than account balance")
	}

	c.AvailableBalance += amount

	return nil
}

// PayAmount decrements the account balance due to payment
func (c *Card) PayAmount(amount float64) error {
	if c.AccountBalance-amount < c.AvailableBalance {
		return errors.Errorf("account balance cannot be lower than available balance")
	}

	c.AccountBalance -= amount

	return nil
}

// BlockAmount decrements the available balance since some amount is blocked
// due to an authorization request
func (c *Card) BlockAmount(amount float64) error {
	if c.AvailableBalance < amount {
		return errors.Errorf("insufficient available balance")
	}

	c.AvailableBalance -= amount

	return nil
}
