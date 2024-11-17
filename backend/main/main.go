package main

import (
	"adopt-pethub/backend/config"
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/logging"

	"adopt-pethub/backend/api"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	//Start config
	c := config.MustParseConfig()
	// Initializes logging
	log.Info("setting up logging")
	logging.InitLogging(c.Log)

	// Initializes the database connection
	db, err := database.NewDatabaseWithMigrations(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	config.SetupCors(e)
	// Initialize the server
	api.RegisterHTTPRoutes(e, db)
	e.Logger.Fatal(e.Start(":5802"))

}
