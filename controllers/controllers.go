package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mvgv/rpgservices/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListarPersonagens(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personagens)
}
