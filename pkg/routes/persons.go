package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigurePersonRoutes(r *mux.Router, controllers *controllers.PersonController) {
	router := r.PathPrefix("/persons").Subrouter()

	//	Create
	router.HandleFunc("", controllers.CreatePerson).Methods("POST")

	//	Read
	router.HandleFunc("", controllers.GetAllPersons).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/{uid:%s}", UUIDRegex), controllers.GetPersonByUID).Methods("GET")

	//	Update
	router.HandleFunc(fmt.Sprintf("/{uid:%s}", UUIDRegex), controllers.UpdatePersonByUID).Methods("PUT")

	//	Delete
	router.HandleFunc(fmt.Sprintf("/{uid:%s}", UUIDRegex), controllers.DeletePersonByUID).Methods("DELETE")
}