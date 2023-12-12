package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureUserNetworkRoutes(r *mux.Router, controllers *controllers.UserNetworkController){
	router := r.PathPrefix("/user_networks").Subrouter()

	//	CREATE
	router.HandleFunc("", controllers.CreateUserNetwork).Methods("POST")

	//	READ
	router.HandleFunc("", controllers.GetAllUserNetworks).Methods("GET")

	//	UPDATE
	router.HandleFunc("/{id:[0-9]+}", controllers.UpdateUserNetworkByID).Methods("PATCH")

	//	DELETE
	router.HandleFunc("/{id:[0-9]+}", controllers.DeleteUserNetworkByID).Methods("DELETE")
}