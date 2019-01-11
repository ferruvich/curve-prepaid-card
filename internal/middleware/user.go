package middleware

import (
	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=user_mock.go -source=user.go -package=middleware -self_package=. User

// User represents the user middleware interface
type User interface {
	Create() (*model.User, error)
	Read(string) (*model.User, error)
}

// UserMiddleware is the User implementation
type UserMiddleware struct {
	middleware Middleware
}

// Create creates and returns new users, or an error if there is one
func (u *UserMiddleware) Create() (*model.User, error) {

	user, err := model.NewUser()
	if err != nil {
		return nil, err
	}

	if err = u.middleware.DataBase().User().Write(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Read returns an existing user
func (u *UserMiddleware) Read(userID string) (*model.User, error) {

	user, err := u.middleware.DataBase().User().Read(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
