package repository

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type RepositoryAnimais struct {
	Logger *logrus.Logger
}

// Interface do repositório para animais
type RepositoryAnimaisInterface interface {
	GetAnimais(db *database.Database) ([]domain.Animal, error)
	InsertAnimal(animal domain.Animal, db *database.Database) error
}

// Implementação do método GetAnimais
func (r *RepositoryAnimais) GetAnimais(db *database.Database) ([]domain.Animal, error) {
	var animais []domain.Animal
	err := db.Connection.Find(&animais).Error
	if err != nil {
		return nil, errors.New("failed to get animals from database")
	}
	return animais, nil
}

func (r *RepositoryAnimais) InsertAnimal(animal domain.Animal, db *database.Database) error {
	if err := db.Connection.Create(&animal).Error; err != nil {
		r.Logger.WithFields(logrus.Fields{
			"animal": animal,
			"error":  err.Error(),
		}).Error("Failed to insert animal into database")
		return fmt.Errorf("failed to insert animal into database: %w", err)
	}

	return nil
}
