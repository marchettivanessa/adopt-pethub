package api

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/handler"
	"adopt-pethub/backend/repository"
	"github.com/labstack/echo"
)

// Registrar rotas para animais
func RegisterAnimalRoutes(e *echo.Echo, db *database.Database) {
	animalRepo := &repository.RepositoryAnimais{}
	animalHandler := handler.NewAnimalHandler(animalRepo)

	e.GET("/animais", func(c echo.Context) error {
		return animalHandler.GetAnimais(c, db)
	})

	e.POST("/animais", func(c echo.Context) error {
		return animalHandler.InsertAnimal(c, db)
	})
}
