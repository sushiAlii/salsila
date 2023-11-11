package models

import "errors"

var (
	ErrUserNotFound				= 	errors.New("User not found")
	ErrRoleIDRequired			= 	errors.New("Role ID is required")
	ErrEmailRequired			= 	errors.New("Email is required")
	ErrEmailNotUnique			= 	errors.New("Email is already in use")
	ErrPasswordRequired 		= 	errors.New("Password is required")
	ErrPasswordMinChar 			= 	errors.New("Password should have more than 4 characters")
	ErrPasswordIncorrect 		= 	errors.New("Incorrect Password!")

	ErrFirstNameRequired		=	errors.New("First name is required")
	ErrMiddleNameRequired		= 	errors.New("Middle name is required")
	ErrLastNameRequired			=	errors.New("Last name is required")
	ErrBirthdayInvalid			=	errors.New("Invalid birthday")
	ErrGenderRequired			=	errors.New("Gender is required")
	ErrGenderInvalid			=	errors.New("Gender should only be on specified options")

	ErrUserUIDRequired			=	errors.New("User UID is required")
	ErrSocialNetworkIDRequired 	= 	errors.New("Social Network ID is required")
	ErrUserURLRequired			=	errors.New("User URL is required")
	ErrUserURLMinChar			=	errors.New("User URL should have more than 5 characters")
)