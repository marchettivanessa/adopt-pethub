package domain

import (
	"adopt-pethub/backend/database"
	"time"
)

type Usuario struct {
	ID           int        `gorm:"primaryKey"`
	Nome         string     `gorm:"column:nome"`
	Email        string     `gorm:"column:email"`
	Senha        string     `gorm:"column:senha"`
	Telefone     *string    `gorm:"column:telefone"`
	Endereco     *string    `gorm:"column:endereco"`
	TipoUsuario  *string    `gorm:"column:tipo_usuario"`
	DataCadastro time.Time  `gorm:"column:data_cadastro"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

type usuario struct{}

// Usuarios define the contract methods for this entity domain
type Usuarios interface {
	GetUsuarioById(id int, db *database.Database) (*Usuario, error)
}

// Creates a new instance of Usuario
func NewUsuario() Usuarios {
	return usuario{}
}

// Método que busca um usuário pelo ID usando o GORM
func (u usuario) GetUsuarioById(id int, db *database.Database) (*Usuario, error) {
	var user Usuario
	if err := db.Connection.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
