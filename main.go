package main

import (
	"fmt"

	"github.com/mvgv/rpgservices/database"
	"github.com/mvgv/rpgservices/models"
	"github.com/mvgv/rpgservices/routes"
)

func main() {
	database.DbConnection()
	fmt.Println("Iniciando o go")
	models.Personagens = []models.Personagem{
		{Id: 1, Nome: "Tragdush", Classe: "Bárbaro"},
		{Id: 2, Nome: "Salazar", Classe: "Mago"},
	}
	routes.HandleRequest()
}
