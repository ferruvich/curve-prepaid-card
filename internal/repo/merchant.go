package repo

import (
	"context"
	"database/sql"

	// Mandatory for PSQL
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/configuration"
	"github.com/ferruvich/curve-challenge/pkg/psql"
)

//go:generate mockgen -destination=merchant_mock.go -package=repo github.com/ferruvich/curve-challenge/internal/repo Merchant

// Merchant is the interface that contains all DB function for merchant
type Merchant interface {
	Write(context.Context, *model.Merchant) error
}

// MerchantRepo handler merchant operations in DB
type MerchantRepo struct {
	dbConnection *sql.DB
}

// NewMerchantRepo initialize the db connection and
// returns the initialized structure
func NewMerchantRepo(ctx context.Context) (Merchant, error) {

	cfg, ok := ctx.Value("cfg").(*configuration.Configuration)
	if !ok {
		return nil, errors.Errorf("error loading configuration")
	}

	db, err := sql.Open(cfg.Psql.DriverName, newSessionString(*cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &MerchantRepo{
		dbConnection: db,
	}, nil
}

// Write writes a new merchant on DB
func (mr *MerchantRepo) Write(ctx context.Context, merchant *model.Merchant) error {

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt("INSERT INTO merchants VALUES ($1)", merchant.ID),
	}

	err := psql.WithTransaction(mr.dbConnection, func(tx psql.Transaction) error {
		_, err := psql.RunPipeline(tx, statements...)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "error writing merchant")
	}

	return nil
}
