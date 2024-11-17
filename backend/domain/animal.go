package domain

import "time"

type Animal struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	Nome         string     `json:"nome" gorm:"column:nome"`
	Especie      string     `json:"especie,omitempty" gorm:"column:especie"`
	Raca         string     `json:"raca,omitempty" gorm:"column:raca"`
	Idade        int        `json:"idade,omitempty" gorm:"column:idade"`
	Sexo         string     `json:"sexo" gorm:"column:sexo"`
	Vacinado     bool       `json:"vacinado" gorm:"column:vacinado"`
	Vermifugado  bool       `json:"vermifugado" gorm:"column:vermifugado"`
	Castrado     bool       `json:"castrado" gorm:"column:castrado"`
	Descricao    string     `json:"descricao" gorm:"column:descricao"`
	FotoURL      string     `json:"foto_url,omitempty" gorm:"column:foto_url"`
	StatusAdocao string     `json:"status_adocao" gorm:"column:status_adocao"`
	DataResgate  *time.Time `json:"data_resgate,omitempty" gorm:"column:data_resgate"`
	DataCadastro *time.Time `json:"data_cadastro,omitempty" gorm:"column:data_cadastro"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type RepositoryAnimal interface{}

func (Animal) TableName() string {
	return "adopt_pethub.animais"
}
