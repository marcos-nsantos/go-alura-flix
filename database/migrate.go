package database

import (
	"github.com/marcos-nsantos/alura-flix/models"
	"log"
)

func Migrate() {
	db, err := Connect()
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&models.Video{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&models.Categoria{})
	if err != nil {
		log.Panic("Erro ao criar tabela de categorias: ", err)
	}

	if err := db.AutoMigrate(models.User{}); err != nil {
		log.Panic("Erro ao criar tabela de usuários: ", err)
	}
}
