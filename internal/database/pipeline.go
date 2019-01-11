package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

//go:generate mockgen -destination=pipeline_mock.go -source=pipeline.go -package=database -self_package=. Pipeline

// Pipeline is a simple wrapper for creating a statement
type Pipeline interface {
	makeQuery(Transaction) (*sql.Rows, error)
}

// pipelineStmt is the pipeline struct
// consisting of a query and a set of arguments to be passed to that query.
type pipelineStmt struct {
	query string
	args  []interface{}
}

// runPipeline runs the supplied statements within the Transaction. If any statement fails, the Transaction
// is rolled back, and the original error is returned.
func (s *Service) runPipeline(tx Transaction, stmts ...*pipelineStmt) (*sql.Rows, error) {
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

// makeQuery Executes the statement within supplied Transaction, returning rows
func (ps *pipelineStmt) makeQuery(tx Transaction) (*sql.Rows, error) {
	return tx.Query(ps.query, ps.args...)
}
