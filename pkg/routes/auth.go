package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	//	Register
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")

	//	Login
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	//	Logout
	router.HandleFunc("/logout", controllers.LogoutUser).Methods("POST")

	//	Refresh token
	router.HandleFunc("/refresh", controllers.RefreshToken).Methods("POST")
}