package repo

import (
	"context"
	"database/sql"

	// Mandatory for PSQL
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-challenge/api/model"
	"github.com/ferruvich/curve-challenge/internal/configuration"
	"github.com/ferruvich/curve-challenge/pkg/psql"
)

//go:generate mockgen -destination=user_mock.go -package=repo github.com/ferruvich/curve-challenge/internal/repo User

// User is the interface that contains all DB function for user
type User interface {
	Write(context.Context, *model.User) error
}

// UserRepo handler user write operation in DB
type UserRepo struct {
	dbConnection *sql.DB
}

// NewUserRepo initialize the db connection and
// returns the initialized structure
func NewUserRepo(ctx context.Context) (User, error) {

	cfg, ok := ctx.Value("cfg").(*configuration.Configuration)
	if !ok {
		return nil, errors.Errorf("error loading configuration")
	}

	db, err := sql.Open(cfg.Psql.DriverName, newSessionString(*cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &UserRepo{
		dbConnection: db,
	}, nil
}

// Write writes a new user on DB
func (ur *UserRepo) Write(ctx context.Context, user *model.User) error {

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt("INSERT INTO users VALUES ($1)", user.ID),
	}

	err := psql.WithTransaction(ur.dbConnection, func(tx psql.Transaction) error {
		_, err := psql.RunPipeline(tx, statements...)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "error writing user")
	}

	return nil
}
