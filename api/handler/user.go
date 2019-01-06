package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-challenge/api/middleware"
)

// User represents the User handler
type User struct {
	middleware middleware.User
}

// NewUserHandler returns a newly created user handler
func NewUserHandler() (Handler, error) {
	middleware, err := middleware.NewUserMiddleware()
	if err != nil {
		return nil, err
	}

	return &User{middleware}, nil
}

// Create is the HTTP handler of the POST /user
func (u *User) Create(c *gin.Context) {

	user, err := u.middleware.Create()
	if err != nil {
		fmt.Printf("%+v", err)
		c.JSON(http.StatusInternalServerError, ErrorMessage{
			Error: "create user failed",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
