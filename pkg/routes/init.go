package routes

import (
	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
)

func ConfigureAllRoutes(r *mux.Router, c controllers.Controllers) {
	ConfigureRoleRoutes(r, c.RoleController)
	ConfigureSocialNetworkRoutes(r, c.SocialNetworkController)
	ConfigureFamilyRoutes(r, c.FamilyController)
	ConfigureUserRoutes(r, c.UserController)
	ConfigureAuthRoutes(r, c.AuthController)
	ConfigurePersonRoutes(r, c.PersonController)
	ConfigureUserNetworkRoutes(r, c.UserNetworkController)
}