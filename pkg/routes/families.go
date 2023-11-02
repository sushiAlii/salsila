package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureFamilyRoutes(r *mux.Router) {
	router := r.PathPrefix("/families").Subrouter()

	router.HandleFunc("", controllers.CreateFamily).Methods("POST")
	router.HandleFunc("", controllers.GetAllFamilies).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.GetFamilyByID).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.UpdateFamilyByID).Methods("PATCH")
	router.HandleFunc("/{id:[0-9]+}", controllers.DeleteFamilyByID).Methods("DELETE")
}