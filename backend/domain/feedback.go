package domain

type Feedback struct {
	ID            int     `gorm:"primaryKey"`
	UsuarioID     int     `gorm:"column:usuario_id"`
	AnimalID      int     `gorm:"column:animal_id"`
	Mensagem      string  `gorm:"column:mensagem"`
	DataAvaliacao string  `gorm:"column:data_avaliacao"`
	Avaliacao     int     `gorm:"column:avaliacao"`
	UpdatedAt     *string `gorm:"column:updated_at, omitempty"`
	DeletedAt     *string `gorm:"column:deleted_at, omitempty"`
}

type RepositoryFeedback interface {
	
}