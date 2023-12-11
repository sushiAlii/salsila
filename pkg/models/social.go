package models

import (
	"gorm.io/gorm"
)

type SocialNetwork struct {
	ID			int				`gorm:"primaryKey" json:"id"`
	Name		string			`gorm:"not null" json:"name"`
	BaseUrl		string			`json:"baseUrl"`
}

type SocialNetworkService interface {
	CreateSocialNetwork(*SocialNetwork) error
	GetAllSocialNetworks() ([]SocialNetwork, error)
	GetSocialNetworkByID(uint) (*SocialNetwork, error)
	UpdateSocialNetworkByID(uint, SocialNetwork) error
	DeleteSocialNetworkByID(uint) error
}

type socialNetworkService struct {
	DB *gorm.DB
}

func NewSocialNetworkService(db *gorm.DB) SocialNetworkService {
	return &socialNetworkService{DB: db}
}

func (sns *socialNetworkService) CreateSocialNetwork (socialNetwork *SocialNetwork) error {
	return sns.DB.Create(socialNetwork).Error
}

func (sns *socialNetworkService) GetAllSocialNetworks() ([]SocialNetwork, error) {
	var socialNetworksList []SocialNetwork

	if err := sns.DB.Find(&socialNetworksList).Error; err != nil {
		return nil, err
	}

	return socialNetworksList, nil
}

func (sns *socialNetworkService) GetSocialNetworkByID(id uint) (*SocialNetwork, error) {
	var socialNetworkData SocialNetwork

	if err := sns.DB.First(&socialNetworkData, id).Error; err != nil {
		return nil, err
	}

	return &socialNetworkData, nil
}

func (sns *socialNetworkService) UpdateSocialNetworkByID(id uint, updatedSocialNetwork SocialNetwork) error {
	return sns.DB.Model(&SocialNetwork{}).Where("id = ?", id).Updates(updatedSocialNetwork).Error
}

func (sns *socialNetworkService) DeleteSocialNetworkByID(id uint) error {
	return sns.DB.Delete(&SocialNetwork{}, id).Error
}