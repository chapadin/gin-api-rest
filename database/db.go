package database

import (
	"Victor/gin-api-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		panic("Erro ao conecta com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
