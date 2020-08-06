package main

import (
	"gqlapi/config"
	"gqlapi/database"
	"gqlapi/logging"
	"gqlapi/schema"
	"gqlapi/server"
	"log"
)

func main() {
	// Config
	conf, err := config.ReadConfig("./config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	logger, err := logging.NewLogger(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	// Database
	db := database.NewDatabase(conf, logger)
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Schema
	schema := schema.NewSchema(db, logger)

	// Server
	serv := server.NewServer(db, conf, logger, schema)
	log.Fatal(serv.Run())
}
