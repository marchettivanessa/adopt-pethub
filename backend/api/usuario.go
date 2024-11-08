package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/middleware"
	"adopt-pethub/backend/repository"

	"github.com/labstack/echo"
)

func RegisterUsuarioRoutes(e *echo.Echo, db *database.Database) {
	usuarioRepo := repository.UsuarioRepository{}
	usuarioHandler := handler.NewHandler(usuarioRepo)

	// Adiciona o banco de dados ao contexto, se necessário
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Insere o db no contexto
			c.Set("db", db)
			return next(c)
		}
	})

	e.GET("/usuario/:id", usuarioHandler.GetUsuarioById, middleware.AuthMiddleware)
	e.GET("/usuarios", usuarioHandler.GetUsuarios, middleware.AuthMiddleware)
}
