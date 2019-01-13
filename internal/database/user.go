package database

import (
	"database/sql"

	// Mandatory for PSQL
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
)

//go:generate mockgen -destination=user_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database User

// User is the interface that contains all DB function for user
type User interface {
	Write(*model.User) error
	Read(string) (*model.User, error)
}

// UserDataBase handler user write operation in DB
type UserDataBase struct {
	service DataBase
}

// Write writes a new user on DB
func (u *UserDataBase) Write(user *model.User) error {

	statements := []*pipelineStmt{
		u.service.newPipelineStmt("INSERT INTO users VALUES ($1)", user.ID),
	}

	_, err := u.service.withTransaction(u.service.GetConnection(), func(tx DBTransaction) (*sql.Rows, error) {
		_, err := u.service.runPipeline(tx, statements...)
		return nil, err
	})
	if err != nil {
		return errors.Wrap(err, "error writing user")
	}

	return nil
}

// Read reds a user from DB
func (u *UserDataBase) Read(userID string) (*model.User, error) {

	user := &model.User{}

	statements := []*pipelineStmt{
		u.service.newPipelineStmt("SELECT * FROM users WHERE ID=$1", userID),
	}

	_, err := u.service.withTransaction(u.service.GetConnection(), func(tx DBTransaction) (*sql.Rows, error) {
		res, err := u.service.runPipeline(tx, statements...)
		if !res.Next() {
			return nil, errors.Errorf("user not found")
		}
		if err = res.Scan(&user.ID); err != nil {
			return nil, errors.Wrap(err, "error building user struct")
		}
		return res, err
	})
	if err != nil {
		return nil, errors.Wrap(err, "error reading user")
	}

	return user, nil
}
