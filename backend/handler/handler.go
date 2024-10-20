package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Handler struct {
	dbMethods dbInterface
}

func NewHandler(dbMethods dbInterface) Handler {
	return Handler{
		dbMethods: dbMethods,
	}
}

// Definindo a interface que contém o método GetUsuarioById
type dbInterface interface {
	GetUsuarioById(id int, db *database.Database) (*domain.Usuario, error)
}

// Handler que busca um usuário pelo ID
func (h Handler) GetUsuarioById(c echo.Context, db *database.Database) error {
	// Obtendo o ID da URL e convertendo para int
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID format")
	}

	// Chamando o método para buscar o usuário pelo ID
	user, err := h.dbMethods.GetUsuarioById(id, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	// Retornando o usuário em formato JSON
	return c.JSON(http.StatusOK, user)
}
