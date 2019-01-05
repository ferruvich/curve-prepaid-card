package middleware

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-challenge/api/model"
)

// User represents the user middleware interface
type User interface {
	Create() (*model.User, error)
}

// UserMiddleware is the User implementation
type UserMiddleware struct{}

// NewUserMiddleware returns a new middleware for user
func NewUserMiddleware() User {
	return &UserMiddleware{}
}

// Create creates and returns new users, or an error if there is one
func (u *UserMiddleware) Create() (*model.User, error) {

	userUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating user id")
	}

	return &model.User{
		ID: userUUID.String(),
	}, nil
}
