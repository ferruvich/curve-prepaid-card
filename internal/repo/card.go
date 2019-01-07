package repo

import (
	context "context"
	"database/sql"

	model "github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/configuration"
	"github.com/ferruvich/curve-challenge/pkg/psql"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=card_mock.go -package=repo github.com/ferruvich/curve-challenge/internal/repo Card

// Card is the interface that contains all DB function for cards
type Card interface {
	Write(context.Context, *model.Card) error
}

// CardRepo handler card operations in DB
type CardRepo struct {
	dbConnection *sql.DB
}

// NewCardRepo initialize the db connection and
// returns the initialized structure
func NewCardRepo(ctx context.Context) (Card, error) {

	cfg, ok := ctx.Value("cfg").(*configuration.Configuration)
	if !ok {
		return nil, errors.Errorf("error loading configuration")
	}

	db, err := sql.Open(cfg.Psql.DriverName, newSessionString(*cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &CardRepo{
		dbConnection: db,
	}, nil
}

// Write writes a new card on DB
func (c *CardRepo) Write(ctx context.Context, card *model.Card) error {
	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt(
			"INSERT INTO cards(ID,owner,account_balance,blocked_amount) VALUES ($1,$2,$3,$4)",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := psql.WithTransaction(c.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		_, err := psql.RunPipeline(tx, statements...)
		return nil, err
	})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
