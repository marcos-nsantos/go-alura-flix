package main

import (
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/routes"
	"github.com/marcos-nsantos/alura-flix/utils"
)

func init() {
	database.Migrate()
	utils.RegisterValidators()
}

func main() {
	routes.HandleRequests()
}
