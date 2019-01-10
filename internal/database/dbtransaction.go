package database

import (
	"database/sql"
)

//go:generate mockgen -destination=dbtransaction_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database Transaction

// Transaction is an interface that models the standard Transaction in
// `database/sql`.
// To ensure funcs in withTransaction cannot commit or rollback a Transaction (which is
// handled by `withTransaction`), those methods are not included here.
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// withTransaction creates a new Transaction and handles rollback/commit based on the
// error object returned by the function fn, and returns its response
func (s *Service) withTransaction(db *sql.DB, fn func(Transaction) (*sql.Rows, error)) (*sql.Rows, error) {
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
