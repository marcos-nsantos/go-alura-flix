package main

import (
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/routes"
	"github.com/marcos-nsantos/alura-flix/utils"
	"log"
)

func init() {
	database.Migrate()
	utils.RegisterValidators()
}

func main() {
	r := routes.HandleRequests()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
