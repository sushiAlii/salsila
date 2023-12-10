package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureRoleRoutes(r *mux.Router, controllers *controllers.RoleController) {
	roleRouter := r.PathPrefix("/roles").Subrouter()

	roleRouter.HandleFunc("", controllers.GetAllRoles).Methods("GET")
	roleRouter.HandleFunc("/{id:[0-9]+}", controllers.GetRoleByID).Methods("GET")
	roleRouter.HandleFunc("", controllers.CreateRole).Methods("POST")
	roleRouter.HandleFunc("/{id:[0-9]+}", controllers.UpdateRoleByID).Methods("PATCH")
	roleRouter.HandleFunc("/{id:[0-9]+}", controllers.DeleteRoleByID).Methods("DELETE")
}