package api

import (
	"adopt-pethub/backend/database"

	"github.com/labstack/echo/v4"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	RegisterUsuarioRoutes(e, db)
	RegisterAnimalRoutes(e, db)
	RegisterFeedbackRoutes(e, db)
	// RegisterAbrigoRoutes(e, db)
}
