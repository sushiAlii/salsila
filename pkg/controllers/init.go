package controllers

import (
	"github.com/sushiAlii/salsila/pkg/models"
)

type Controllers struct {
	RoleController 			*RoleController
	SocialNetworkController *SocialNetworkController
	FamilyController		*FamilyController
	UserController			*UserController
	AuthController			*AuthController
	PersonController		*PersonController
	UserNetworkController	*UserNetworkController
}

func InstantiateControllers(s models.Services) Controllers {
	roleController := NewRoleController(s.RoleService)

	socialNetworkController := NewSocialNetworkController(s.SocialNetworkService)

	familyController := NewFamilyController(s.FamilyService)

	userController := NewUserController(s.UserService)

	authController := NewAuthController(s.AuthService, s.UserService)

	personController := NewPersonController(s.PersonService)

	userNetworkController := NewUserNetworkController(s.UserNetworkService)

	return Controllers{
		RoleController: roleController,
		SocialNetworkController: socialNetworkController,
		FamilyController: familyController,
		UserController: userController,
		AuthController: authController,
		PersonController: personController,
		UserNetworkController: userNetworkController,
	}
}