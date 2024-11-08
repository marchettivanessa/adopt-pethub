package repository

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
	"errors"
)

type RepositoryAnimais struct{}

// Interface do repositório para animais
type RepositoryAnimaisInterface interface {
	GetAnimais(db *database.Database) ([]domain.Animal, error)
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
