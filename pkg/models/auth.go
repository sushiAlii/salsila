package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(DB *gorm.DB, email, password string) (*User, error) {
	var user User

	_, err := GetUsersByEmail(DB, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrPasswordIncorrect
	}

	return &user, nil
}

func RegisterUser(DB *gorm.DB, user *User) error {
	return CreateUser(DB, user)
}