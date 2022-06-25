package main

import (
	"fmt"

	"github.com/felipejzsouza/go-api-rest/database"
	"github.com/felipejzsouza/go-api-rest/models"
	"github.com/felipejzsouza/go-api-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Felipe", Historia: "Estudante"},
		{Id: 2, Nome: "Nubia", Historia: "Recrutadora"},
	}
	database.ConectaDB()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
