package main

import (
	"adopt-pethub/backend/config"
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/logging"

	"adopt-pethub/backend/api"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	//Start config
	c := config.MustParseConfig()
	// Initializes logging.
	log.Info("setting up logging")
	logging.InitLogging(c.Log)

	// Iniciating database
	// Initializes the database connection.
	db, err := database.NewDatabaseWithMigrations(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Inicie o servidor
	e := echo.New()
	api.RegisterHTTPRoutes(e, db)
	e.Logger.Fatal(e.Start(":5802"))

}
