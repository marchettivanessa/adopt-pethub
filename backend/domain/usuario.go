package domain

import (
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

type RepositoryUsuario interface {}

func (Usuario) TableName() string {
    return "adopt_pethub.usuarios"
}