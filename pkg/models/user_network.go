package models

import (
	"time"

	"gorm.io/gorm"
)


type UserNetwork struct {
	ID				uint		`gorm:"primaryKey" json:"id"`
	UserUID			string		`gorm:"not null" json:"userUid"`
	SocialNetworkID	uint		`gorm:"not null" json:"socialNetworkId"`
	UserName		string		`json:"userName"`
	UserURL			string		`gorm:"not null" json:"userUrl"`
	CreatedAt		time.Time	`gorm:"type:timestamptz" json:"-"`
	UpdatedAt		*time.Time	`gorm:"type:timestamptz" json:"-"`
}

func ValidateUserNetwork(DB *gorm.DB, userNetwork *UserNetwork) error {
	if userNetwork.UserUID == "" {
		return ErrUserUIDRequired
	}

	if userNetwork.SocialNetworkID == 0 {
		return ErrSocialNetworkIDRequired
	}

	if userNetwork.UserURL == "" {
		return ErrUserURLRequired
	}

	if len(userNetwork.UserURL) < 5 {
		return ErrUserURLMinChar
	}

	return nil
}


func CreateUserNetwork(DB *gorm.DB, userNetwork *UserNetwork) error {
	return DB.Create(userNetwork).Error
}

func GetAllUserNetworks(DB *gorm.DB) ([]UserNetwork, error) {
	var userNetworks []UserNetwork

	if err := DB.Find(&userNetworks).Error; err != nil {
		return nil, err
	}

	return userNetworks, nil
}

func GetUserNetworksByUserUID(DB *gorm.DB, userUid string) ([]UserNetwork, error) {
	var userNetworks []UserNetwork

	if err := DB.Find(&userNetworks, "user_uid = ?", userUid).Error; err != nil {
		return nil, err
	}

	return userNetworks, nil
}

func GetUserNetworkByID(DB *gorm.DB, id uint) (*UserNetwork, error) {
	var userNetwork UserNetwork

	if err := DB.First(&userNetwork, id).Error; err != nil {
		return nil, err
	}

	return &userNetwork, nil
}

func UpdateUserNetworkByID(DB *gorm.DB, userNetwork *UserNetwork, id uint) error {
	tx := DB.Begin()

	if err := tx.Model(&UserNetwork{}).Where("id = ?", id).Updates(&userNetwork).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func DeleteUserNetworkByID(DB *gorm.DB, id uint) error {
	tx := DB.Begin()

	if err := tx.Delete(&UserNetwork{}, id).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}