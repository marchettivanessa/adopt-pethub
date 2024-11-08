package use_cases

import "adopt-pethub/backend/domain"

type listAnimals struct{}
type ListAnimals interface {
	Execute() ([]domain.Animal, error)
}

func NewListAnimals() *listAnimals {
	return &listAnimals{}
}

func (u *listAnimals) Execute() ([]domain.Animal, error) {
// buscar todos os animais - chamar banco de dados
// checar se o animal ainda está disponível
	animals := []domain.Animal{}
	return animals, nil
}