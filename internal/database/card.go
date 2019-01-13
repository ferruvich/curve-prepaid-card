package database

import (
	"database/sql"

	"github.com/ferruvich/curve-prepaid-card/internal/model"

	"github.com/pkg/errors"
)

//go:generate mockgen -destination=card_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database Card

// Card is the interface that contains all DB function for cards
type Card interface {
	Write(*model.Card) error
	Update(*model.Card) error
	Read(string) (*model.Card, error)
}

// CardDataBase handles card operations in DB
type CardDataBase struct {
	service DataBase
}

// Write writes a new card on DB
func (c *CardDataBase) Write(card *model.Card) error {

	statements := []*pipelineStmt{
		c.service.newPipelineStmt(
			"INSERT INTO cards(ID,owner,account_balance,blocked_amount) VALUES ($1,$2,$3,$4)",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := c.service.withTransaction(c.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			_, err := c.service.runPipeline(tx, statements...)
			return nil, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}

// Read reds a card from DB
func (c *CardDataBase) Read(cardID string) (*model.Card, error) {

	updatedCard := &model.Card{}
	blockedAmount := 0.0

	statements := []*pipelineStmt{
		c.service.newPipelineStmt("SELECT * FROM cards WHERE ID=$1", cardID),
	}

	_, err := c.service.withTransaction(c.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			res, err := c.service.runPipeline(tx, statements...)
			if !res.Next() {
				return nil, errors.Errorf("card not found")
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
func (c *CardDataBase) Update(card *model.Card) error {

	statements := []*pipelineStmt{
		c.service.newPipelineStmt(
			"UPDATE cards SET owner=$2, account_balance=$3, blocked_amount=$4 where ID=$1",
			card.ID, card.Owner, card.AccountBalance, (card.AccountBalance - card.AvailableBalance),
		),
	}

	_, err := c.service.withTransaction(c.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			res, err := c.service.runPipeline(tx, statements...)
			return res, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
