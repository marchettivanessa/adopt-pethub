// api/router.go
package api

import (
	"adopt-pethub/backend/database"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	// Chamando a função que registra as rotas de cada domínio
	RegisterUsuarioRoutes(e, db)
	RegisterAnimalRoutes(e, db)
	// RegisterFeedbackRoutes(e, db)
}
