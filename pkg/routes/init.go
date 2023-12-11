package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAllRoutes(r *mux.Router, roleController *controllers.RoleController, socialNetworkController *controllers.SocialNetworkController, familyController *controllers.FamilyController) {
	ConfigureRoleRoutes(r, roleController)
	ConfigureSocialNetworkRoutes(r, socialNetworkController)
	ConfigureFamilyRoutes(r, familyController)
	ConfigureAuthRoutes(r)
	ConfigureUserRoutes(r)
	ConfigurePersonRoutes(r)
	ConfigureUserNetworkRoutes(r)
}