package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// User represents an User inside system
type User struct {
	ID string `json:"ID"`
}

func NewUser() (*User, error) {

	userUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.Wrap(err, "error generating user id")
	}

	return &User{
		ID: userUUID.String(),
	}, nil
}
