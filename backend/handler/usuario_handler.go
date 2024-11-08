package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (h UsuarioHandler) GetUsuarioById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID format")
	}

	// Obtendo o banco de dados diretamente do contexto
	db := c.Get("db").(*database.Database)

	user, err := h.usuarioDbMethods.GetUsuarioById(id, db)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (h UsuarioHandler) GetUsuarios(c echo.Context) error {
	db := c.Get("db").(*database.Database)

	users, err := h.usuarioDbMethods.GetUsuarios(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get users")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UsuarioHandler) RegisterUsuario(c echo.Context, db *database.Database) error {
	var newUser domain.Usuario
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
	}
	// TODO finish this method
	// Hashing da senha (bcrypt, por exemplo)
	// Salvar o novo usuário no banco

	// Após salvar, enviar a resposta
	return c.JSON(http.StatusCreated, newUser)
}

func (h *UsuarioHandler) Login(c echo.Context, db *database.Database) error {
	var loginUser struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	// Bind request body data
	if err := c.Bind(&loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
	}

	var usuario domain.Usuario
	if err := db.Connection.Where("email = ?", loginUser.Email).First(&usuario).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	// Compare and validate password

	// Generates token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": usuario.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
