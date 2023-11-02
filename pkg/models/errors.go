package models

import "errors"

var (
	ErrRoleIDRequired	= errors.New("Role ID is required")
	ErrEmailRequired	= errors.New("Email is required")
	ErrEmailNotUnique	= errors.New("Email is already in use")
	ErrPasswordRequired = errors.New("Password is required")
	ErrPasswordMinChar 	= errors.New("Password should have more than 4 characters")
)