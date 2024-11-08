package handler

import (
	"adopt-pethub/backend/database"
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
