package main

import (
	"log"

	"github.com/OlegDjur/Go-Auth-Microservice/config"
	"github.com/OlegDjur/Go-Auth-Microservice/pkg/db"
)

func main() {
	cfgFile, err := config.LoadConfig("./config/config-local")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	psqlDB, err := db.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}
	defer psqlDB.Close()
}
