package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/curve-prepaid-card/internal/middleware"
)

// User represents the User handler
type User interface {
	Create() func(c *gin.Context)
}

// UserHandler is the User struct
type UserHandler struct {
	server Server
}

// Create is the HTTP handler of the POST /user
func (u *UserHandler) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := middleware.NewMiddleware(u.server.DataBase()).User().Create()
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage{
				Error: "create user failed",
			})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
