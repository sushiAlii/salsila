package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureSocialNetworkRoutes(r *mux.Router) {
	router := r.PathPrefix("/social_networks").Subrouter()

	router.HandleFunc("", controllers.GetAllSocialNetworks).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.GetSocialNetworkByID).Methods("GET")

	router.HandleFunc("", controllers.CreateSocialNetwork).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}", controllers.UpdateSocialNetworkByID).Methods("PATCH")

	router.HandleFunc("/{id:[0-9]+}", controllers.DeleteSocialNetworkByID).Methods("DELETE")
}