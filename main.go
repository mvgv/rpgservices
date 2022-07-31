package main

import (
	"fmt"

	"github.com/mvgv/rpgservices/models"
	"github.com/mvgv/rpgservices/routes"
)

func main() {
	fmt.Println("Iniciando o go")
	models.Personagens = []models.Personagem{
		{Nome: "Tragdush", Classe: "Bárbaro"},
		{Nome: "Salazar", Classe: "Mago"},
	}
	routes.HandleRequest()
}
