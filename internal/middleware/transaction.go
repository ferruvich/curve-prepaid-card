package middleware

import (
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -source=transaction.go -destination=transaction_mock.go -package=middleware -self_package=. Transaction

// Transaction is the transaction middleware interface
type Transaction interface {
	CreatePayment(string, float64) (*model.Transaction, error)
}

// TransactionMiddleware is the Transaction implementation
type TransactionMiddleware struct {
	middleware Middleware
}

// CreatePayment creates new payment transaction
func (t *TransactionMiddleware) CreatePayment(authReqID string, amount float64) (*model.Transaction, error) {

	authReq, err := t.middleware.DataBase().AuthorizationRequest().Read(authReqID)
	if err != nil {
		return nil, err
	}

	if err = authReq.Capture(amount); err != nil {
		return nil, err
	}

	card, err := t.middleware.DataBase().Card().Read(authReq.Card)
	if err != nil {
		return nil, err
	}

	if err = card.PayAmount(amount); err != nil {
		return nil, err
	}

	tx, err := model.NewPaymentTransaction(card.Owner, authReq.Merchant, amount)
	if err != nil {
		return nil, err
	}

	if err = t.middleware.DataBase().AuthorizationRequest().Update(authReq); err != nil {
		return nil, err
	}

	if err = t.middleware.DataBase().Card().Update(card); err != nil {
		return nil, err
	}

	if err = t.middleware.DataBase().Transaction().Write(tx); err != nil {
		return nil, err
	}

	return tx, nil
}
