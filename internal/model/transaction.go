package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	payment = "payment"
	refund  = "refund"

	dateFormat = "2006-01-02"
)

// Transaction embeds an exchange of money
type Transaction struct {
	ID       string  `json:"ID"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
	Date     string  `json:"date"`
	Type     string  `json:"type"`
}

// NewTransaction returns a new transaction with id
func NewTransaction(sender, receiver string, amount float64) (*Transaction, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating id")
	}

	tx := &Transaction{ID: id.String()}

	tx.Sender = sender
	tx.Receiver = receiver
	tx.Amount = amount
	tx.Date = time.Now().Format(dateFormat)

	return tx, nil
}

// NewPaymentTransaction returns a new transaction
func NewPaymentTransaction(sender, receiver string, amount float64) (*Transaction, error) {
	tx, err := NewTransaction(sender, receiver, amount)
	if err != nil {
		return nil, err
	}

	tx.Type = payment

	return tx, nil
}

// NewRefundTransaction returns a new refund
func NewRefundTransaction(sender, receiver string, amount float64) (*Transaction, error) {
	tx, err := NewTransaction(sender, receiver, amount)
	if err != nil {
		return nil, err
	}

	tx.Type = refund

	return tx, nil
}
