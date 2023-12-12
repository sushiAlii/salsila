package models

import (
	"time"

	"gorm.io/gorm"
)

type UserNetwork struct {
	ID					uint		`gorm:"primaryKey" json:"id"`
	UserUID				string		`gorm:"not null" json:"userUid"`
	SocialNetworksID	uint		`gorm:"not null" json:"socialNetworkId"`
	UserName			string		`json:"userName"`
	UserURL				string		`gorm:"not null" json:"userUrl"`
	CreatedAt			time.Time	`gorm:"type:timestamptz" json:"-"`
	UpdatedAt			*time.Time	`gorm:"type:timestamptz" json:"-"`
}

type UserNetworkService interface {
	CreateUserNetwork(*UserNetwork) error
	GetAllUserNetworks() ([]UserNetwork, error)
	GetUserNetworksByUserUID(string) ([]UserNetwork, error)
	GetUserNetworkByID(uint) (*UserNetwork, error)
	UpdateUserNetworkByID(*UserNetwork, uint) error
	DeleteUserNetworkByID(uint) error
}

type userNetworkService struct {
	DB *gorm.DB
}

func NewUserNetworkService(db *gorm.DB) UserNetworkService {
	return &userNetworkService{DB: db}
}

func ValidateCreateUserNetwork(userNetwork *UserNetwork) error {
	if userNetwork.UserUID == "" {
		return ErrUserUIDRequired
	}

	if userNetwork.SocialNetworksID == 0 {
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

func ValidateUpdateUserNetwork(userNetwork *UserNetwork) error {
	if userNetwork.UserURL == "" {
		return ErrUserURLRequired
	}

	if len(userNetwork.UserURL) < 5 {
		return ErrUserURLMinChar
	}

	return nil
}

func (uns *userNetworkService) CreateUserNetwork(userNetwork *UserNetwork) error {
	return uns.DB.Create(userNetwork).Error
}

func (uns *userNetworkService) GetAllUserNetworks() ([]UserNetwork, error) {
	var userNetworks []UserNetwork

	if err := uns.DB.Find(&userNetworks).Error; err != nil {
		return nil, err
	}

	return userNetworks, nil
}

func (uns *userNetworkService) GetUserNetworksByUserUID(userUid string) ([]UserNetwork, error) {
	var userNetworks []UserNetwork

	if err := uns.DB.Find(&userNetworks, "user_uid = ?", userUid).Error; err != nil {
		return nil, err
	}

	return userNetworks, nil
}

func (uns *userNetworkService) GetUserNetworkByID(id uint) (*UserNetwork, error) {
	var userNetwork UserNetwork

	if err := uns.DB.First(&userNetwork, id).Error; err != nil {
		return nil, err
	}

	return &userNetwork, nil
}

func (uns *userNetworkService) UpdateUserNetworkByID(userNetwork *UserNetwork, id uint) error {
	tx := uns.DB.Begin()

	if err := tx.Model(&UserNetwork{}).Where("id = ?", id).Updates(&userNetwork).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func (uns *userNetworkService) DeleteUserNetworkByID(id uint) error {
	tx := uns.DB.Begin()

	if err := tx.Delete(&UserNetwork{}, id).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}