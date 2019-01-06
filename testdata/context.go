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
		GinMode:     "debug",
		Environment: "development",
		Server: &configuration.Server{
			Host: "localhost",
			Port: "8080",
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
