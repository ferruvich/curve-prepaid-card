package middleware

import "github.com/ferruvich/curve-prepaid-card/internal/database"

// Middleware is the middlewares main interface
type Middleware interface {
	DataBase() database.DataBase
	AuthorizationRequest() AuthorizationRequest
	Card() Card
	Merchant() Merchant
	User() User
}

// Service is the Middleware struct
type Service struct {
	db database.DataBase
}

// NewMiddleware returns a new middleware service
func NewMiddleware(db database.DataBase) Middleware {
	return &Service{
		db: db,
	}
}

// DataBase returns the database from middleware
func (s *Service) DataBase() database.DataBase {
	return s.db
}

// AuthorizationRequest returns a new AuthorizationRequest middleware
func (s *Service) AuthorizationRequest() AuthorizationRequest {
	return &AuthorizationRequestMiddleware{}
}

// Card returns a new Card middleware
func (s *Service) Card() Card {
	return &CardMiddleware{}
}

// Merchant returns a new Merchant middleware
func (s *Service) Merchant() Merchant {
	return &MerchantMiddleware{}
}

// User returns a new User middleware
func (s *Service) User() User {
	return &UserMiddleware{}
}
