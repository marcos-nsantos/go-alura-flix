package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() (*gorm.DB, error) {
	dns := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return db, nil
}
