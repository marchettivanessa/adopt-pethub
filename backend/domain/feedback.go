package domain

import "time"

type Feedback struct {
	ID            int        `gorm:"primaryKey" json:"id"`
	UsuarioID     int        `gorm:"column:usuario_id" json:"usuario_id,omitempty"`
	AnimalID      *int       `gorm:"column:animal_id" json:"animal_id,omitempty"`
	Mensagem      string     `gorm:"column:mensagem" json:"mensagem"`
	DataAvaliacao time.Time  `gorm:"column:data_avaliacao" json:"data_avaliacao"`
	Avaliacao     int        `gorm:"column:avaliacao" json:"avaliacao"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

type RepositoryFeedback interface{}

func (Feedback) TableName() string {
	return "adopt_pethub.feedbacks"
}
