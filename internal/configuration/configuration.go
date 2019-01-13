package configuration

import (
	"os"
)

// Configuration embeds full configuration
type Configuration struct {
	Server *Server
	Psql   *Psql
}

// Server embeds server configuration
type Server struct {
	Host string
	Port string
}

// Psql embeds Postgresql configuration
type Psql struct {
	DriverName string
	DBName     string
	User       string
	Host       string
	SSLMode    string
}

// GetConfiguration loads configuration from JSON file
// and returns it, or it returns an error
func GetConfiguration() *Configuration {

	return &Configuration{
		Server: &Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Psql: &Psql{
			DriverName: os.Getenv("PSQL_DRIVER_NAME"),
			DBName:     os.Getenv("PSQL_DB_NAME"),
			User:       os.Getenv("PSQL_USER"),
			Host:       os.Getenv("PSQL_HOST"),
			SSLMode:    os.Getenv("PSQL_SSL_MODE"),
		},
	}
}
