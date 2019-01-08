package testdata

import (
	"context"
	"testing"

	"github.com/ferruvich/curve-challenge/internal/configuration"
)

// GetMockContext returns prepared context for tests
func GetMockContext(t *testing.T) context.Context {

	t.Helper()

	conf := &configuration.Configuration{
		Server: &configuration.Server{
			GinMode: "debug",
			Host:    "localhost",
			Port:    "8080",
		},
		Psql: &configuration.Psql{
			DriverName: "postgres",
			DBName:     "curve",
			User:       "postgres",
			Host:       "localhost",
			SSLMode:    "disable",
		},
	}

	return context.WithValue(context.Background(), "cfg", conf)
}
