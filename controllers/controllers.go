package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mvgv/rpgservices/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListarPersonagens(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personagens)
}

func DetalharPersonagem(w http.ResponseWriter, r *http.Request) {
	idPersonagem := mux.Vars(r)["id"]
	idNumerico, err := strconv.Atoi(idPersonagem)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(models.Personagens[idNumerico-1])
}
