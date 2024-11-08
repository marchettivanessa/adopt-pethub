package repository

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"
)

type UsuarioRepository struct{}

func (u UsuarioRepository) GetUsuarioById(id int, db *database.Database) (*domain.Usuario, error) {
	var user domain.Usuario
	if err := db.Connection.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UsuarioRepository) GetUsuarios(db *database.Database) ([]domain.Usuario, error) {
	var users []domain.Usuario
	if err := db.Connection.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
