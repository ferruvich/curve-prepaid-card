package database

import (
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=authorization_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database AuthorizationRequest

// AuthorizationRequest is the interface that contains all DB function for cards
type AuthorizationRequest interface {
	Write(*model.AuthorizationRequest) error
}

// AuthorizationRequestDataBase handler card operations in DB
type AuthorizationRequestDataBase struct {
	service DataBase
}

// Write writes a new card on DB
func (a *AuthorizationRequestDataBase) Write(authReq *model.AuthorizationRequest) error {

	statements := []*pipelineStmt{
		a.service.newPipelineStmt(
			"INSERT INTO authorizations(ID,merchant,card,approved,amount,reversed) VALUES ($1,$2,$3,$4,$5,$6)",
			authReq.ID, authReq.Merchant, authReq.Card, authReq.Approved, authReq.Amount, authReq.Reversed,
		),
	}

	_, err := a.service.withTransaction(a.service.GetConnection(),
		func(tx Transaction) (*sql.Rows, error) {
			_, err := a.service.runPipeline(tx, statements...)
			return nil, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
