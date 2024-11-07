package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UsuarioHandler struct {
	usuarioDbMethods usuarioInterface
}

func NewHandler(usuarioDB usuarioInterface) UsuarioHandler {
	return UsuarioHandler{
		usuarioDbMethods: usuarioDB,
	}
}

// Definindo a interface que contém o método GetUsuarioById
type usuarioInterface interface {
	GetUsuarioById(id int, db *database.Database) (*domain.Usuario, error)
	GetUsuarios(*database.Database) ([]domain.Usuario, error)
}

// Handler que busca um usuário pelo ID
func (h UsuarioHandler) GetUsuarioById(c echo.Context, db *database.Database) error {
	// Obtendo o ID da URL e convertendo para int
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID format")
	}

	user, err := h.usuarioDbMethods.GetUsuarioById(id, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, user)
}


// Handler que busca todos os usuários
func (h UsuarioHandler) GetUsuarios(c echo.Context, db *database.Database) error {
	users, err := h.usuarioDbMethods.GetUsuarios(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get users")
	}
	return c.JSON(http.StatusOK, users)
}
