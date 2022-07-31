package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mvgv/rpgservices/database"
	"github.com/mvgv/rpgservices/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListarPersonagens(w http.ResponseWriter, r *http.Request) {
	var listaPersonagens []models.Personagem
	database.DB.Find(&listaPersonagens)
	json.NewEncoder(w).Encode(listaPersonagens)
}

func DetalharPersonagem(w http.ResponseWriter, r *http.Request) {
	idPersonagem := mux.Vars(r)["id"]
	idPersonagemNumerico, err := strconv.Atoi(idPersonagem)
	if err != nil {
		log.Fatal(err)
	}
	var personagem models.Personagem
	database.DB.First(&personagem, idPersonagemNumerico)
	json.NewEncoder(w).Encode(personagem)
}
