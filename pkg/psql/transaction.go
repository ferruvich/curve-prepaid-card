package psql

import (
	"database/sql"
)

//go:generate mockgen -destination=transaction_mock.go -package=psql github.com/ferruvich/curve-challenge/pkg/psql Transaction

// Transaction is an interface that models the standard transaction in
// `database/sql`.
// To ensure funcs in WithTransaction cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the function fn, and returns its response
func WithTransaction(db *sql.DB, fn func(Transaction) (*sql.Rows, error)) (*sql.Rows, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	return fn(tx)
}
