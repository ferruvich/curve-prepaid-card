package database

import (
	"context"
	"database/sql"

	// Mandatory for PSQL
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
	"github.com/ferruvich/curve-prepaid-card/internal/psql"
)

//go:generate mockgen -destination=user_mock.go -package=database github.com/ferruvich/curve-prepaid-card/internal/database User

// User is the interface that contains all DB function for user
type User interface {
	Write(context.Context, *model.User) error
	Read(context.Context, string) (*model.User, error)
}

// Userdatabase handler user write operation in DB
type Userdatabase struct {
	dbConnection *sql.DB
}

// NewUserdatabase initialize the db connection and
// returns the initialized structure
func NewUserdatabase(ctx context.Context) (User, error) {

	cfg, ok := ctx.Value("cfg").(*configuration.Configuration)
	if !ok {
		return nil, errors.Errorf("error loading configuration")
	}

	db, err := sql.Open(cfg.Psql.DriverName, newSessionString(*cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &Userdatabase{
		dbConnection: db,
	}, nil
}

// Write writes a new user on DB
func (ur *Userdatabase) Write(ctx context.Context, user *model.User) error {

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt("INSERT INTO users VALUES ($1)", user.ID),
	}

	_, err := psql.WithTransaction(ur.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		_, err := psql.RunPipeline(tx, statements...)
		return nil, err
	})
	if err != nil {
		return errors.Wrap(err, "error writing user")
	}

	return nil
}

// Read reds a user from DB
func (ur *Userdatabase) Read(ctx context.Context, userID string) (*model.User, error) {

	user := &model.User{}

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt("SELECT * FROM users WHERE ID=$1", userID),
	}

	_, err := psql.WithTransaction(ur.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		res, err := psql.RunPipeline(tx, statements...)
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
