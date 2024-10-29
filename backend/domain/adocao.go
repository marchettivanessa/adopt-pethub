package domain

type Adocao struct {
	ID           int     `gorm:"primaryKey"`
	IDAnimal     int     `gorm:"column:id_animal"`
	IDUsuario    int     `gorm:"column:id_usuario"`
	DataAdocao   string  `gorm:"column:data_adocao"`
	StatusAdocao *string `gorm:"column:status_adocao"`
	Observacoes  *string `gorm:"column:observacoes"`
	UpdatedAt    *string `gorm:"column:updated_at, omitempty"`
	DeletedAt    *string `gorm:"column:deleted_at, omitempty"`
}

type RepositoryAdocao interface {}