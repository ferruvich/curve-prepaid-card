package repo

import "fmt"

const (
	sessionString = "host=%s user=%s dbname=%s sslmode=%s"

	host    = "localhost"
	user    = "postgres"
	dbName  = "curve"
	sslMode = "disable"
)

// newSessionString returns the session string needed in sql.Open()
func newSessionString() string {
	return fmt.Sprintf(sessionString, host, user, dbName, sslMode)
}
