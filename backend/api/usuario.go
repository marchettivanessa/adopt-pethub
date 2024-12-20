package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/middleware"
	"adopt-pethub/backend/repository"

	"github.com/labstack/echo/v4"
)

func RegisterUsuarioRoutes(e *echo.Echo, db *database.Database) {
	usuarioRepo := repository.UsuarioRepository{}
	usuarioHandler := handler.NewHandler(usuarioRepo)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Insere o db no contexto
			c.Set("db", db)
			return next(c)
		}
	})

	e.GET("/usuario/:id", usuarioHandler.GetUsuarioById, middleware.AuthMiddleware)
	e.GET("/usuarios", usuarioHandler.GetUsuarios, middleware.AuthMiddleware)
	e.POST("/login", usuarioHandler.Login)
	e.POST("/usuario", usuarioHandler.CreateUsuario)

}
