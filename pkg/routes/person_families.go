package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigurePersonFamilyRoutes(r *mux.Router) {
	router := r.PathPrefix("/person_families").Subrouter()

	//	CREATE
	router.HandleFunc("", controllers.CreatePersonFamily).Methods("POST")

	//	READ
	router.HandleFunc("", controllers.GetAllPersonsFamilies).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.GetPersonFamilyByID).Methods("GET")
	
	//	UPDATE
	router.HandleFunc("/{id:[0-9]+}", controllers.UpdatePersonFamilyByID).Methods("PUT")

	//	DELETE
	router.HandleFunc("/{id:[0-9]+}", controllers.DeletePersonsFamilyByID).Methods("DELETE")
}