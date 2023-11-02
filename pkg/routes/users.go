package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

const UUIDRegex = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"

func ConfigureUserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/{uid:%s}", UUIDRegex), controllers.GetUserByUID).Methods("GET")
}