package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAllRoutes(r *mux.Router, roleController *controllers.RoleController, socialNetworkController *controllers.SocialNetworkController, familyController *controllers.FamilyController, authController *controllers.AuthController) {
	ConfigureRoleRoutes(r, roleController)
	ConfigureSocialNetworkRoutes(r, socialNetworkController)
	ConfigureFamilyRoutes(r, familyController)
	ConfigureAuthRoutes(r, authController)
	ConfigureUserRoutes(r)
	ConfigurePersonRoutes(r)
	ConfigureUserNetworkRoutes(r)
}