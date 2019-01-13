package database

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -source=transaction.go -destination=transaction_mock.go -package=database -self_package=. Transaction

// Transaction is the interface that contains all DB functions for transactions
type Transaction interface {
	Write(*model.Transaction) error
}

// TransactionDataBase is the Transaction struct
type TransactionDataBase struct {
	service DataBase
}

func (t *TransactionDataBase) Write(tx *model.Transaction) error {

	statements := []*pipelineStmt{
		t.service.newPipelineStmt(
			"INSERT INTO transactions(ID,sender,receiver,amount,date,type) VALUES ($1,$2,$2,$4,$5,$6)",
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
