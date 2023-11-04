package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UID 		string 			`gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"uid"`
	RoleID		uint			`gorm:"not null" json:"roleId"`
	PersonsUID	*string			`json:"personsUid"`
	Email 		string			`gorm:"uniqueIndex;not null" json:"email"`
	Password	string			`gorm:"not null" json:"password,omitempty"`
	CreatedAt	time.Time		`gorm:"type:timestamptz" json:"-"`
	UpdatedAt	*time.Time		`gorm:"type:timestamptz" json:"-"`
	DeletedAt	gorm.DeletedAt	`gorm:"type:timestamptz" json:"-"`
}


func ValidateUser(DB *gorm.DB, user *User) error {
	if user.RoleID == 0 {
		return ErrRoleIDRequired
	}

	if strings.TrimSpace(user.Email) == "" {
		return ErrEmailRequired
	}

	var existingUser User

	fmt.Printf("Finding email with %s", user.Email)

	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrEmailNotUnique
		}
	}

	password := strings.TrimSpace(user.Password)

	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) <= 4 {
		return ErrPasswordMinChar
	}

	return nil
}

func CreateUser(DB *gorm.DB, newUser *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)

	return DB.Omit("uid").Create(newUser).Error
}

func GetAllUsers(DB *gorm.DB) ([]User, error) {
	var usersList []User

	if err := DB.Select("uid, role_id, persons_uid, email").Find(&usersList).Error; err != nil {
		return nil, err
	}

	return usersList, nil
}

func GetUserByUID(DB *gorm.DB, uid string) (*User, error) {
	var user User

	if err := DB.Select("uid, role_id, persons_uid, email").Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsersByEmail(DB *gorm.DB, email string) ([]User, error) {
	var users []User

	if err := DB.Select("uid, role_id, persons_uid, email").
				Where("email LIKE ?", "%" + email + "%").
					Find(&users).Error; err != nil {
			
		return nil, err
	}

	return users, nil
}

func DeleteUserByUID(DB *gorm.DB, uid string) error {
	return DB.Delete(&User{}, "uid = ?", uid).Error
}