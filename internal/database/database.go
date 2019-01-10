package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const (
	sessionTemplate = "host=%s user=%s dbname=%s sslmode=%s"
)

//go:generate mockgen -destination=database_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database DataBase

// DataBase represents the entry point for the DB
type DataBase interface {
	withTransaction(*sql.DB, func(Transaction) (*sql.Rows, error)) (*sql.Rows, error)
	runPipeline(Transaction, ...*pipelineStmt) (*sql.Rows, error)
	newPipelineStmt(string, ...interface{}) *pipelineStmt
	AuthorizationRequest() AuthorizationRequest
	Card() Card
	Merchant() Merchant
	User() User
}

// Service is the DataBase struct
type Service struct {
	dbConnection *sql.DB
}

// NewDatabaseService returns a new DB service
func NewDatabaseService(driverName, host, user, dbName, sslMode string) (DataBase, error) {

	session := fmt.Sprintf(
		sessionTemplate, host, user, dbName, sslMode,
	)

	db, err := sql.Open(driverName, session)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &Service{dbConnection: db}, nil
}

// AuthorizationRequest returns interface for user operations on DB
func (s *Service) AuthorizationRequest() AuthorizationRequest {
	return &AuthorizationRequestDataBase{
		service: s,
	}
}

// Card returns interface for card operations on DB
func (s *Service) Card() Card {
	return &CardDataBase{
		service: s,
	}
}

// Merchant returns interface for merchant operations on DB
func (s *Service) Merchant() Merchant {
	return &MerchantDataBase{
		service: s,
	}
}

// User returns interface for user operations on DB
func (s *Service) User() User {
	return &UserDataBase{
		service: s,
	}
}
