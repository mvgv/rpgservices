package routes

import (
	"log"
	"net/http"

	"github.com/mvgv/rpgservices/controllers"
)

func HandleRequest() {

	http.HandleFunc("/rpgs/poa/personagens", controllers.ListarPersonagens)
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
