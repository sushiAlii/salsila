package models

import (
	"gorm.io/gorm"
)

type Services struct {
	RoleService 			RoleService
	SocialNetworkService 	SocialNetworkService
	FamilyService			FamilyService
	UserService				UserService
	AuthService				AuthService
	PersonService			PersonService
	UserNetworkService		UserNetworkService
}

func InstantiateServices(dbInstance *gorm.DB) Services {
	roleService := NewRoleService(dbInstance)

	socialNetworkService := NewSocialNetworkService(dbInstance)

	familyService := NewFamilyController(dbInstance)

	userService := NewUserService(dbInstance)

	authService := NewAuthService(dbInstance, userService)

	personService := NewPersonService(dbInstance)

	userNetworkService := NewUserNetworkService(dbInstance)

	return Services{
		RoleService: roleService, 
		SocialNetworkService: socialNetworkService, 
		FamilyService: familyService,
		UserService: userService,
		AuthService: authService,
		PersonService: personService,
		UserNetworkService: userNetworkService,
	}
}