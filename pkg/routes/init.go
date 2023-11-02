package routes

import "github.com/gorilla/mux"

func ConfigureAllRoutes(r *mux.Router) {
	ConfigureRoleRoutes(r)
	ConfigureSocialNetworkRoutes(r)
	ConfigureFamilyRoutes(r)
	ConfigureAuthRoutes(r)
	ConfigureUserRoutes(r)
}