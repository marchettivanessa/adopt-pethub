package handler

import (
	"adopt-pethub/backend/config"
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioHandler struct {
	usuarioDbMethods Repository
}

func NewHandler(usuarioDB Repository) UsuarioHandler {
	return UsuarioHandler{
		usuarioDbMethods: usuarioDB,
	}
}

type Repository interface {
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

func (h *UsuarioHandler) CreateUsuario(c echo.Context) error {
	var newUser domain.Usuario
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err))
	}

	hashedPassword, err := Hash(newUser.Senha)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to hash password")
	}
	newUser.Senha = string(hashedPassword)

	db := c.Get("db").(*database.Database)
	if err := db.Connection.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to register user")
	}

	// Remove password from response
	newUser.Senha = ""
	return c.JSON(http.StatusCreated, newUser)
}

func (h *UsuarioHandler) Login(c echo.Context) error {
	var loginUser struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}
	var usuario domain.Usuario

	// Bind faz a leitura e validação direta do corpo JSON
	if err := c.Bind(&loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	db := c.Get("db").(*database.Database)
	if err := db.Connection.Where("email = ?", loginUser.Email).First(&usuario).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	if err := h.ComparePassword(loginUser.Senha, usuario.Senha); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	// Gerarates the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": usuario.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.MustParseConfig().JWTSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func Hash(senha string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (h *UsuarioHandler) ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
