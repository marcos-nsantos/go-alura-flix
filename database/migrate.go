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
}
