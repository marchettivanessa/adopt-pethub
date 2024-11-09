package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/middleware"
	"adopt-pethub/backend/repository"

	"github.com/labstack/echo"
)

func RegisterAnimalRoutes(e *echo.Echo, db *database.Database) {
	animalRepo := &repository.RepositoryAnimais{}
	animalHandler := handler.NewAnimalHandler(animalRepo)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	e.GET("/animais", animalHandler.GetAnimais, middleware.AuthMiddleware)

	e.POST("/animais", animalHandler.InsertAnimal, middleware.AuthMiddleware)
}
