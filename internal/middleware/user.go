package middleware

import (
	"context"

	"github.com/ferruvich/curve-prepaid-card/internal/database"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

// User represents the user middleware interface
type User interface {
	Create(context.Context) (*model.User, error)
	Read(context.Context, string) (*model.User, error)
}

// UserMiddleware is the User implementation
type UserMiddleware struct {
	database database.User
}

// NewUserMiddleware returns a new middleware for user
func NewUserMiddleware(ctx context.Context) (User, error) {
	database, err := database.NewUserdatabase(ctx)
	if err != nil {
		return nil, err
	}

	return &UserMiddleware{
		database: database,
	}, nil
}

// Create creates and returns new users, or an error if there is one
func (u *UserMiddleware) Create(ctx context.Context) (*model.User, error) {

	user, err := model.NewUser()
	if err != nil {
		return nil, err
	}

	if err = u.database.Write(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Read returns an existing user
func (u *UserMiddleware) Read(ctx context.Context, userID string) (*model.User, error) {

	user, err := u.database.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
