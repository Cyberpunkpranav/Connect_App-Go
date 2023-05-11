package routes

import (
	"ConnectApp/src/models"
	"fmt"

	"github.com/gorilla/mux"
)

var VaccineList = func(router *mux.Router) {
	router.HandleFunc("/Vaccines/List", models.GetVaccines).Methods("GET")
}

func main() {
	fmt.Println("")
}
