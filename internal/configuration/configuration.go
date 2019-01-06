package configuration

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"github.com/tkanos/gonfig"
)

const (
	// FILENAME is the configuration file name
	FILENAME = "configuration.json"
	// PRODUCTIONENV is used to recognise if we are in production or not
	PRODUCTIONENV = "production"
	// DBDOCKERHOST is the DB docker container name
	DBDOCKERHOST = "db"
)

// Configuration embeds full configuration
type Configuration struct {
	Environment string  `json:"ENVIRONMENT"`
	GinMode     string  `json:"GIN_MODE"`
	Server      *Server `json:"SERVER"`
	Psql        *Psql   `json:"PSQL"`
}

// Server embeds server configuration
type Server struct {
	Host string `json:"HOST"`
	Port string `json:"PORT"`
}

// Psql embeds Postgresql configuration
type Psql struct {
	DriverName string `json:"DRIVER"`
	DBName     string `json:"DBNAME"`
	User       string `json:"USER"`
	Host       string `json:"HOST"`
	SSLMode    string `json:"SSL"`
}

// GetConfiguration loads configuration from JSON file
// and returns it, or it returns an error
func GetConfiguration() (*Configuration, error) {

	configuration := &Configuration{}

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), FILENAME)

	if err := gonfig.GetConf(filePath, configuration); err != nil {
		return nil, errors.Wrapf(err, "Error loading configuration file")
	}

	if configuration.Environment == PRODUCTIONENV {
		configuration.Psql.Host = DBDOCKERHOST
	}

	return configuration, nil
}
