package domain

type Animal struct {
	ID           int     `gorm:"primaryKey"`
	Nome         string  `gorm:"column:nome"`
	Especie      string  `gorm:"column:especie"`
	Raca         *string `gorm:"column:raca, omitempty"`
	Idade        int     `gorm:"column:idade"`
	Sexo         *string `gorm:"column:sexo"`
	Vacinado     bool    `gorm:"column:vacinado"`
	Vermifugado  bool    `gorm:"column:vermifugado"`
	Castrado     bool    `gorm:"column:castrado"`
	Descricao    string  `gorm:"column:descricao"`
	FotoUrl      string  `gorm:"column:foto_url, omitempty"`
	StatusAdocao *string `gorm:"column:status_adocao"`
	DataResgate  string  `gorm:"column:data_resgate"`
	DataCadastro string  `gorm:"column:data_cadastro"`
	UpdatedAt    *string `gorm:"column:updated_at, omitempty"`
	DeletedAt    *string `gorm:"column:deleted_at, omitempty"`
}

type RepositoryAnimal interface {}