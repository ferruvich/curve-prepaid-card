package middleware

import (
	"github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/repo"
)

// User represents the user middleware interface
type User interface {
	Create() (*model.User, error)
}

// UserMiddleware is the User implementation
type UserMiddleware struct {
	repo repo.User
}

// NewUserMiddleware returns a new middleware for user
func NewUserMiddleware() (User, error) {
	repo, err := repo.NewUserRepo()
	if err != nil {
		return nil, err
	}

	return &UserMiddleware{
		repo: repo,
	}, nil
}

// Create creates and returns new users, or an error if there is one
func (u *UserMiddleware) Create() (*model.User, error) {

	user, err := model.NewUser()
	if err != nil {
		return nil, err
	}

	if err = u.repo.Write(user); err != nil {
		return nil, err
	}

	return user, nil
}
