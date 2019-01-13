package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/database"
)

// Server is the main interface
type Server interface {
	Routers() *gin.Engine
	DataBase() database.DataBase
	NewAuthorizationRequestHandler() AuthorizationRequest
	NewCardHandler() Card
	NewMerchantHandler() Merchant
	NewTransactionHandler() Transaction
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

// Routers set ups and return a gin router
func (s *Service) Routers() *gin.Engine {
	router := gin.Default()

	// Routes
	router.POST("/user", s.NewUserHandler().Create())
	router.POST("/user/:userID/card", s.NewCardHandler().Create())
	router.GET("/user/:userID/card/:cardID", s.NewCardHandler().GetCard())
	router.POST("/user/:userID/card/:cardID/deposit", s.NewCardHandler().Deposit())
	router.POST("/merchant", s.NewMerchantHandler().Create())
	router.POST("/authorization", s.NewAuthorizationRequestHandler().Create())
	router.POST("/authorization/:authID/capture", s.NewTransactionHandler().Create())

	return router
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

// NewTransactionHandler returns a new Transaction handler
func (s *Service) NewTransactionHandler() Transaction {
	return &TransactionHandler{
		server: s,
	}
}

// NewUserHandler returns a new User handler
func (s *Service) NewUserHandler() User {
	return &UserHandler{
		server: s,
	}
}
