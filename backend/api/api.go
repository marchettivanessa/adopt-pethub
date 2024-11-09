// api/router.go
package api

import (
	"adopt-pethub/backend/database"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	RegisterUsuarioRoutes(e, db)
	RegisterAnimalRoutes(e, db)
	RegisterFeedbackRoutes(e, db)
	// RegisterAbrigoRoutes(e, db)
}
