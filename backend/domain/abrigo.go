package domain

type Abrigo struct {
	ID        int     `gorm:"primaryKey"`
	Nome      string  `gorm:"column:nome"`
	Endereco  *string `gorm:"column:endereco"`
	Telefone  *string `gorm:"column:telefone"`
	Email     *string `gorm:"column:email"`
	IdTutor   int     `gorm:"column:id_tutor"`
	UpdatedAt *string `gorm:"column:updated_at, omitempty"`
	DeletedAt *string `gorm:"column:deleted_at, omitempty"`
}

type RepositoryAbrigo interface {}

func (Abrigo) TableName() string {
    return "adopt_pethub.abrigos"
}
