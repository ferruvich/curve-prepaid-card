package database

import (
	"database/sql"

	"github.com/ferruvich/curve-prepaid-card/internal/model"

	"github.com/pkg/errors"
)

//go:generate mockgen -destination=card_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database Card

// Card is the interface that contains all DB function for cards
type Card interface {
	Write(*sql.DB, *model.Card) error
	Update(*sql.DB, *model.Card) error
	Read(*sql.DB, string) (*model.Card, error)
}

// CardDataBase handler card operations in DB
type CardDataBase struct {
	service DataBase
}

// Write writes a new card on DB
func (c *CardDataBase) Write(dbConnection *sql.DB, card *model.Card) error {

	statements := []*pipelineStmt{
		c.service.newPipelineStmt(
			"INSERT INTO cards(ID,owner,account_balance,blocked_amount) VALUES ($1,$2,$3,$4)",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := c.service.withTransaction(dbConnection,
		func(tx transaction) (*sql.Rows, error) {
			_, err := c.service.runPipeline(tx, statements[1])
			return nil, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}

// Read reds a card from DB
func (c *CardDataBase) Read(dbConnection *sql.DB, cardID string) (*model.Card, error) {

	updatedCard := &model.Card{}
	blockedAmount := 0.0

	statements := []*pipelineStmt{
		c.service.newPipelineStmt("SELECT * FROM cards WHERE ID=$1", cardID),
	}

	_, err := c.service.withTransaction(dbConnection,
		func(tx transaction) (*sql.Rows, error) {
			res, err := c.service.runPipeline(tx, statements...)
			if !res.Next() {
				return nil, errors.Errorf("user not found")
			}
			if err = res.Scan(&updatedCard.ID, &updatedCard.Owner,
				&updatedCard.AccountBalance, &blockedAmount); err != nil {
				return nil, errors.Wrap(err, "error building card struct")
			}
			updatedCard.AvailableBalance = updatedCard.AccountBalance - blockedAmount
			return res, err
		})
	if err != nil {
		return nil, errors.Wrap(err, "error reading card")
	}

	return updatedCard, nil
}

// Update updates a card in DB
func (c *CardDataBase) Update(dbConnection *sql.DB, card *model.Card) error {

	statements := []*pipelineStmt{
		c.service.newPipelineStmt(
			"UPDATE cards SET owner=$2, account_balance=$3, blocked_amount=$4 where ID=$1",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := c.service.withTransaction(dbConnection,
		func(tx transaction) (*sql.Rows, error) {
			res, err := c.service.runPipeline(tx, statements...)
			return res, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
