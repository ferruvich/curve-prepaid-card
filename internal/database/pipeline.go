package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// A pipelineStmt is a simple wrapper for creating a statement consisting of
// a query and a set of arguments to be passed to that query.
type pipelineStmt struct {
	query string
	args  []interface{}
}

// runPipeline runs the supplied statements within the transaction. If any statement fails, the transaction
// is rolled back, and the original error is returned.
func (s *Service) runPipeline(tx transaction, stmts ...*pipelineStmt) (*sql.Rows, error) {
	var res *sql.Rows
	var err error

	for _, ps := range stmts {
		res, err = ps.makeQuery(tx)
		if err != nil {
			return nil, errors.Wrapf(err, "Error executing query %q", ps.query)
		}
	}

	return res, nil
}

// newPipelineStmt is used to create PipelineStmt
func (s *Service) newPipelineStmt(query string, args ...interface{}) *pipelineStmt {
	return &pipelineStmt{query, args}
}

// makeQuery Executes the statement within supplied transaction, returning rows
func (ps *pipelineStmt) makeQuery(tx transaction) (*sql.Rows, error) {
	return tx.Query(ps.query, ps.args...)
}
