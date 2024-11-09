package domain

import "time"

type Abrigo struct {
	ID        int        `gorm:"primaryKey;column:id"`
	Nome      string     `gorm:"column:nome"`
	Endereco  *string    `gorm:"column:endereco"`
	Telefone  *string    `gorm:"column:telefone"`
	Email     *string    `gorm:"column:email"`
	IdTutor   int        `gorm:"column:id_tutor"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (Abrigo) TableName() string {
	return "adopt_pethub.abrigos"
}

type RepositoryAbrigo interface{}
