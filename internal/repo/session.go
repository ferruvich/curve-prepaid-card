package repo

import (
	"fmt"

	"github.com/ferruvich/curve-challenge/internal/configuration"
)

// SESSIONSTRING represents the string skeleton of the session
// used to configure db connection
const SESSIONSTRING = "host=%s user=%s dbname=%s sslmode=%s"

// newSessionString returns the session string needed in sql.Open()
func newSessionString(cfg configuration.Configuration) string {
	return fmt.Sprintf(
		SESSIONSTRING, cfg.Psql.Host, cfg.Psql.User, cfg.Psql.DBName, cfg.Psql.SSLMode,
	)
}
