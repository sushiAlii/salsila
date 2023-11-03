package models

import (
	"gorm.io/gorm"
)

type SocialNetwork struct {
	ID			int				`gorm:"primaryKey" json:"id"`
	Name		string			`gorm:"not null" json:"name"`
	BaseUrl		string			`json:"baseUrl"`
}

func CreateSocialNetwork (DB *gorm.DB, socialNetwork *SocialNetwork) error {
	return DB.Create(socialNetwork).Error
}

func GetAllSocialNetworks(DB *gorm.DB) ([]SocialNetwork, error) {
	var socialNetworksList []SocialNetwork

	if err := DB.Find(&socialNetworksList).Error; err != nil {
		return nil, err
	}

	return socialNetworksList, nil
}

func GetSocialNetworkByID(DB *gorm.DB, id uint) (*SocialNetwork, error) {
	var socialNetworkData SocialNetwork

	if err := DB.First(&socialNetworkData, id).Error; err != nil {
		return nil, err
	}

	return &socialNetworkData, nil
}

func UpdateSocialNetworkByID(DB *gorm.DB, id uint, updatedSocialNetwork SocialNetwork) error {
	return DB.Model(&SocialNetwork{}).Where("id = ?", id).Updates(updatedSocialNetwork).Error
}

func DeleteSocialNetworkByID(DB *gorm.DB, id uint) error {
	return DB.Delete(&SocialNetwork{}, id).Error
}