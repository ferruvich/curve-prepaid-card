package psql

import (
	"database/sql"

	"github.com/pkg/errors"
)

// A PipelineStmt is a simple wrapper for creating a statement consisting of
// a query and a set of arguments to be passed to that query.
type PipelineStmt struct {
	query  string
	args   []interface{}
	output interface{}
}

// NewPipelineStmt is used to create PipelineStmt
func NewPipelineStmt(output interface{}, query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args, output}
}

// Exec Executes the statement within supplied transaction.
func (ps *PipelineStmt) Exec(tx Transaction) (sql.Result, error) {
	return tx.Exec(ps.query, ps.args...)
}

// Query Executes the statement within supplied transaction, returning rows
func (ps *PipelineStmt) Query(tx Transaction) (*sql.Rows, error) {
	return tx.Query(ps.query, ps.args...)
}

// RunPipeline runs the supplied statements within the transaction. If any statement fails, the transaction
// is rolled back, and the original error is returned.
func RunPipeline(tx Transaction, stmts ...*PipelineStmt) (interface{}, error) {
	var res interface{}
	var err error

	for _, ps := range stmts {
		if ps.output == nil {
			res, err = ps.Exec(tx)
		} else {
			res, err = ps.Query(tx)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "Error executing query %q", ps.query)
		}
	}

	return res, nil
}
