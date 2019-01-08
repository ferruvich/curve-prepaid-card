package repo

import (
	context "context"
	"database/sql"

	"github.com/ferruvich/curve-challenge/internal/configuration"
	"github.com/ferruvich/curve-challenge/internal/model"
	"github.com/ferruvich/curve-challenge/pkg/psql"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=card_mock.go -package=repo github.com/ferruvich/curve-challenge/internal/repo Card

// Card is the interface that contains all DB function for cards
type Card interface {
	Write(context.Context, *model.Card) error
	Update(context.Context, *model.Card) (*model.Card, error)
	Read(context.Context, string) (*model.Card, error)
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

// Read reds a card from DB
func (c *CardRepo) Read(ctx context.Context, cardID string) (*model.Card, error) {

	updatedCard := &model.Card{}
	blockedAmount := 0.0

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt("SELECT * FROM users WHERE ID=$1", cardID),
	}

	rows, err := psql.WithTransaction(c.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		res, err := psql.RunPipeline(tx, statements...)
		if !res.Next() {
			return nil, errors.Errorf("user not found")
		}
		if err = res.Scan(&updatedCard.ID, &updatedCard.Owner,
			&updatedCard.AccountBalance, &blockedAmount); err != nil {
			return nil, errors.Wrap(err, "error building user struct")
		}
		updatedCard.AvailableBalance = updatedCard.AccountBalance - blockedAmount
		return res, err
	})
	if err != nil {
		return nil, errors.Wrap(err, "error writing user")
	}
	defer rows.Close()

	return updatedCard, nil
}

// Update updates a card in DB
func (c *CardRepo) Update(ctx context.Context, card *model.Card) (*model.Card, error) {

	updatedCard := &model.Card{}
	blockedAmount := 0.0

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt(
			"UPDATE cards SET ID=$1, owner=$2, account_balance=$3, blocked_amount=$4 where ID=$5",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := psql.WithTransaction(c.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		res, err := psql.RunPipeline(tx, statements...)
		if !res.Next() {
			return nil, errors.Errorf("user not found")
		}
		if err = res.Scan(&updatedCard.ID, &updatedCard.Owner,
			&updatedCard.AccountBalance, &blockedAmount); err != nil {
			return nil, errors.Wrap(err, "error building user struct")
		}
		updatedCard.AvailableBalance = updatedCard.AccountBalance - blockedAmount
		return res, err
	})
	if err != nil {
		return nil, errors.Wrap(err, "error writing card")
	}

	return updatedCard, nil
}
