package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"adopt-pethub/backend/handler"

	"github.com/labstack/echo"
)

func RegisterHTTPRoutes(e *echo.Echo, db *database.Database) {
	// Inicializando a lógica de domínio e o handler
	usuarioService := domain.NewUsuario()
	uhandler := handler.NewHandler(usuarioService)

	// Handler para pegar usuário por ID
	getHandler := func(c echo.Context) error {
		return uhandler.GetUsuarioById(c, db)
	}

	// Definindo as rotas
	e.GET("/usuario/:id", getHandler)
}
