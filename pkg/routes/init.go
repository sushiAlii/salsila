package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAllRoutes(r *mux.Router, roleController *controllers.RoleController, socialNetworkController *controllers.SocialNetworkController, familyController *controllers.FamilyController, userController *controllers.UserController, authController *controllers.AuthController, personController *controllers.PersonController, userNetworkController *controllers.UserNetworkController) {
	ConfigureRoleRoutes(r, roleController)
	ConfigureSocialNetworkRoutes(r, socialNetworkController)
	ConfigureFamilyRoutes(r, familyController)
	ConfigureUserRoutes(r, userController)
	ConfigureAuthRoutes(r, authController)
	ConfigurePersonRoutes(r, personController)
	ConfigureUserNetworkRoutes(r, userNetworkController)
}