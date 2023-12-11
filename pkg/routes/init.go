package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAllRoutes(r *mux.Router, roleController *controllers.RoleController, socialNetworkController *controllers.SocialNetworkController) {
	ConfigureRoleRoutes(r, roleController)
	ConfigureSocialNetworkRoutes(r)
	ConfigureFamilyRoutes(r)
	ConfigureAuthRoutes(r)
	ConfigureUserRoutes(r)
	ConfigurePersonRoutes(r)
	ConfigureUserNetworkRoutes(r)
}