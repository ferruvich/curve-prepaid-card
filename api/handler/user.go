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
func NewUserHandler() Handler {
	return &User{
		middleware: middleware.NewUserMiddleware(),
	}
}

// Create is the HTTP handler of the POST /user
func (u *User) Create(c *gin.Context) {

	user, err := u.middleware.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage{
			Error: fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
