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
	Read(string) (*model.AuthorizationRequest, error)
	Update(*model.AuthorizationRequest) error
}

// AuthorizationRequestDataBase handler card operations in DB
type AuthorizationRequestDataBase struct {
	service DataBase
}

// Write writes a new card on DB
func (a *AuthorizationRequestDataBase) Write(authReq *model.AuthorizationRequest) error {

	statements := []*pipelineStmt{
		a.service.newPipelineStmt(
			"INSERT INTO authorizations(ID,merchant,card,approved,amount,reversed,captured) VALUES ($1,$2,$3,$4,$5,$6,$7)",
			authReq.ID, authReq.Merchant, authReq.Card, authReq.Approved,
			authReq.Amount, authReq.Reversed, authReq.Captured,
		),
	}

	_, err := a.service.withTransaction(a.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			_, err := a.service.runPipeline(tx, statements...)
			return nil, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}

// Read reds an auth request from DB
func (a *AuthorizationRequestDataBase) Read(authReqID string) (*model.AuthorizationRequest, error) {

	authReq := &model.AuthorizationRequest{}

	statements := []*pipelineStmt{
		a.service.newPipelineStmt("SELECT * FROM authorizations WHERE ID=$1", authReqID),
	}

	_, err := a.service.withTransaction(a.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			res, err := a.service.runPipeline(tx, statements...)
			if !res.Next() {
				return nil, errors.Errorf("auth request not found")
			}
			if err = res.Scan(&authReq.ID, &authReq.Merchant, &authReq.Card,
				&authReq.Amount, &authReq.Approved, &authReq.Reversed, &authReq.Captured); err != nil {
				return nil, errors.Wrap(err, "error building auth struct")
			}
			return res, err
		})
	if err != nil {
		return nil, errors.Wrap(err, "error reading authorization request")
	}

	return authReq, nil
}

// Update updates a card in DB
func (a *AuthorizationRequestDataBase) Update(authReq *model.AuthorizationRequest) error {

	statements := []*pipelineStmt{
		a.service.newPipelineStmt(
			"UPDATE authorizations SET merchant=$2,card=$3,approved=$4,amount=$5,reversed=$6,captured=$7 WHERE ID=$1;",
			authReq.ID, authReq.Merchant, authReq.Card, authReq.Approved,
			authReq.Amount, authReq.Reversed, authReq.Captured,
		),
	}

	_, err := a.service.withTransaction(a.service.GetConnection(),
		func(tx DBTransaction) (*sql.Rows, error) {
			res, err := a.service.runPipeline(tx, statements...)
			return res, err
		})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
