package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/repository"

	"github.com/labstack/echo"
)

func RegisterUsuarioRoutes(e *echo.Echo, db *database.Database) {
	// Instancia o repositório e o handler de usuário
	usuarioRepo := repository.UsuarioRepository{}
	usuarioHandler := handler.NewHandler(usuarioRepo)

	// Define as rotas do domínio usuário
	e.GET("/usuario/:id", func(c echo.Context) error {
		return usuarioHandler.GetUsuarioById(c, db)
	})
	e.GET("/usuarios", func(c echo.Context) error {
		return usuarioHandler.GetUsuarios(c, db)
	})
}
