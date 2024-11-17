package handler

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"adopt-pethub/backend/logging"
	"adopt-pethub/backend/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AnimalHandler struct {
	repository repository.RepositoryAnimaisInterface
}

func NewAnimalHandler(repository repository.RepositoryAnimaisInterface) *AnimalHandler {
	return &AnimalHandler{repository: repository}
}

func (h *AnimalHandler) GetAnimais(c echo.Context) error {
	db := c.Get("db").(*database.Database)
	animais, err := h.repository.GetAnimais(db)
	if err != nil {
		log := logging.Logger(map[string]interface{}{
			"project":  "adopt-pethub",
			"package":  "handler",
			"function": "GetAnimais",
		})
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, animais)
}

func (h *AnimalHandler) InsertAnimal(c echo.Context) error {
	db, ok := c.Get("db").(*database.Database)
	if !ok || db == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database connection not found"})
	}

	log := logging.Logger(map[string]interface{}{
		"project":  "adopt-pethub",
		"package":  "animal_handler",
		"function": "InsertAnimal",
	})

	if c.Request().ContentLength == 0 {
		log.Error("Nenhum dado recebido na requisição")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nenhum dado recebido"})
	}
	var animal domain.Animal

	idade := c.FormValue("idade")
	idadeInt, err := strconv.Atoi(idade)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid idade value"})
	}

	// Get data from request
	animal.Nome = c.FormValue("nome")
	animal.Especie = c.FormValue("especie")
	animal.Raca = c.FormValue("raca")
	animal.Idade = idadeInt
	animal.Sexo = c.FormValue("sexo")
	animal.Descricao = c.FormValue("descricao")
	animal.Vacinado = c.FormValue("vacinado") == "true"
	animal.Castrado = c.FormValue("castrado") == "true"
	animal.Vermifugado = c.FormValue("vermifugado") == "true"

	// Data de cadastro
	if animal.DataCadastro == nil {
		currentTime := time.Now()
		animal.DataCadastro = &currentTime
	}
	if animal.DataResgate == nil {
		currentTime := time.Now()
		animal.DataResgate = &currentTime
	}

	if animal.StatusAdocao == "" {
		animal.StatusAdocao = "DISPONÍVEL"
	}

	file, err := c.FormFile("foto_url")
	if err == nil {
		// Manipulação de arquivo (opcional: salvar em um diretório)
		src, err := file.Open()
		if err != nil {
			log.Error("failed to open uploaded file")
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file upload"})
		}
		defer src.Close()

		animal.FotoURL = file.Filename
	}

	if err := h.repository.InsertAnimal(animal, db); err != nil {
		log.Error("failed to insert animal into database")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert animal"})
	}

	return c.JSON(http.StatusCreated, animal)
}
