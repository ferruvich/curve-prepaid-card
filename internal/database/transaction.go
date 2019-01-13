package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -source=transaction.go -destination=transaction_mock.go -package=database -self_package=. Transaction

// Transaction is the interface that contains all DB functions for transactions
type Transaction interface {
	Write(*model.Transaction) error
	GetListByCard(string) ([]*model.Transaction, error)
}

// TransactionDataBase is the Transaction struct
type TransactionDataBase struct {
	service DataBase
}

func (t *TransactionDataBase) Write(tx *model.Transaction) error {

	statements := []*pipelineStmt{
		t.service.newPipelineStmt(
			"INSERT INTO transactions(ID,sender,receiver,amount,date,type) VALUES ($1,$2,$3,$4,$5,$6)",
			tx.ID, tx.Sender, tx.Receiver, tx.Amount, tx.Date, tx.Type,
		),
	}

	_, err := t.service.withTransaction(t.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			_, err := t.service.runPipeline(tx, statements...)
			return nil, err
		},
	)
	if err != nil {
		return errors.Wrap(err, "error writing transaction")
	}

	return nil
}

// GetListByCard returns a transaction list, given the card that has been used to make them
func (t *TransactionDataBase) GetListByCard(userID string) ([]*model.Transaction, error) {

	txs := []*model.Transaction{}

	statements := []*pipelineStmt{
		t.service.newPipelineStmt(
			"SELECT * FROM transactions WHERE receiver=$1 OR sender=$1", userID,
		),
	}

	_, err := t.service.withTransaction(t.service.GetConnection(),
		func(dbTx DBTransaction) (*sql.Rows, error) {
			res, err := t.service.runPipeline(dbTx, statements...)
			fmt.Println(res)
			for res.Next() {
				tx := model.Transaction{}
				if err = res.Scan(&tx.ID, &tx.Sender, &tx.Receiver, &tx.Amount, &tx.Date, &tx.Type); err != nil {
					return nil, errors.Wrap(err, "error building card struct")
				}
				txs = append(txs, &tx)
			}
			return res, err
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "error writing transaction")
	}

	return txs, nil
}
