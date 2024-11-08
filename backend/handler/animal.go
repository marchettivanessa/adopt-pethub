package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"adopt-pethub/backend/logging"
	"adopt-pethub/backend/repository"
	"net/http"

	"github.com/labstack/echo"
)

type AnimalHandler struct {
	repository repository.RepositoryAnimaisInterface
}

// Novo handler para animais
func NewAnimalHandler(repository repository.RepositoryAnimaisInterface) *AnimalHandler {
	return &AnimalHandler{repository: repository}
}

// Método do handler para buscar os animais
func (h *AnimalHandler) GetAnimais(c echo.Context, db *database.Database) error {
	animais, err := h.repository.GetAnimais(db)
	if err != nil {
		log := logging.Logger(map[string]interface{} {
			"project": "adopt-pethub",
			"package": "handler",
	})
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, animais)
}

func (h *AnimalHandler) InsertAnimal(c echo.Context, db *database.Database) error {
	log := logging.Logger(map[string]interface{}{
		"project": "adopt-pethub",
		"package": "handler",
	})

	var animal domain.Animal

	// Bind dos dados para a estrutura Animal
	if err := c.Bind(&animal); err != nil {
		log.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Error("failed to bind animal data")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid animal data"})
	}

	if err := h.repository.InsertAnimal(animal, db); err != nil {
		log.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Error("failed to insert animal into database")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert animal"})
	}

	return c.JSON(http.StatusCreated, animal)
}
