package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvgv/rpgservices/controllers"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/rpgs/poa/personagens", controllers.ListarPersonagens).Methods("Get")
	router.HandleFunc("/rpgs/poa/personagens/{id}", controllers.DetalharPersonagem).Methods("Get")
	router.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", router))
}
