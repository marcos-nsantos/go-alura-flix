package main

import (
	"github.com/marcos-nsantos/alura-flix/database"
	"github.com/marcos-nsantos/alura-flix/routes"
)

func init() {
	database.Migrate()
}

func main() {
	routes.HandleRequests()
}
