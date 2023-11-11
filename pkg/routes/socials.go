package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureSocialNetworkRoutes(r *mux.Router) {
	router := r.PathPrefix("/social_networks").Subrouter()

	//	Create
	router.HandleFunc("", controllers.CreateSocialNetwork).Methods("POST")

	//	Read
	router.HandleFunc("", controllers.GetAllSocialNetworks).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controllers.GetSocialNetworkByID).Methods("GET")

	//	Update
	router.HandleFunc("/{id:[0-9]+}", controllers.UpdateSocialNetworkByID).Methods("PATCH")

	//	Delete
	router.HandleFunc("/{id:[0-9]+}", controllers.DeleteSocialNetworkByID).Methods("DELETE")
}