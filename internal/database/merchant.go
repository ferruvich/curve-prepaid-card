package database

import (
	"database/sql"

	// Mandatory for PSQL
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=merchant_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database Merchant

// Merchant is the interface that contains all DB function for merchant
type Merchant interface {
	Write(*sql.DB, *model.Merchant) error
}

// MerchantDataBase handler merchant operations in DB
type MerchantDataBase struct {
	service DataBase
}

// Write writes a new merchant on DB
func (m *MerchantDataBase) Write(dbConnection *sql.DB, merchant *model.Merchant) error {

	statements := []*pipelineStmt{
		m.service.newPipelineStmt("INSERT INTO merchants VALUES ($1)", merchant.ID),
	}

	_, err := m.service.withTransaction(dbConnection, func(tx Transaction) (*sql.Rows, error) {
		_, err := m.service.runPipeline(tx, statements...)
		return nil, err
	})
	if err != nil {
		return errors.Wrap(err, "error writing merchant")
	}

	return nil
}
