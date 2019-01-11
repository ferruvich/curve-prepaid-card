package server

import (
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

// Server is the main interface
type Server interface {
	DataBase() database.DataBase
	NewAuthorizationRequestHandler() AuthorizationRequest
	NewCardHandler() Card
	NewMerchantHandler() Merchant
	NewUserHandler() User
}

// Service is the server struct
type Service struct {
	db database.DataBase
}

// NewServer creates a new Server
func NewServer() (Server, error) {

	conf := configuration.GetConfiguration()

	db, err := database.NewDataBase(
		conf.Psql.DriverName, conf.Psql.Host, conf.Psql.User,
		conf.Psql.DBName, conf.Psql.SSLMode,
	)
	if err != nil {
		return nil, errors.Errorf("error initializing database")
	}

	return &Service{db}, nil
}

// DataBase returns the db instance
func (s *Service) DataBase() database.DataBase {
	return s.db
}

// NewAuthorizationRequestHandler returns a new AuthorizationRequest handler
func (s *Service) NewAuthorizationRequestHandler() AuthorizationRequest {
	return &AuthorizationRequestHandler{
		server: s,
	}
}

// NewCardHandler returns a new Card handler
func (s *Service) NewCardHandler() Card {
	return &CardHandler{
		server: s,
	}
}

// NewMerchantHandler returns a new Merchant handler
func (s *Service) NewMerchantHandler() Merchant {
	return &MerchantHandler{
		server: s,
	}
}

// NewUserHandler returns a new User handler
func (s *Service) NewUserHandler() User {
	return &UserHandler{
		server: s,
	}
}
