package main

import (
	"log"
	"strings"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/server"
)

func main() {
	cfg := configuration.GetConfiguration()

	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	server.Routers().Run(strings.Join(
		[]string{
			cfg.Server.Host,
			cfg.Server.Port,
		}, ":",
	))
}
