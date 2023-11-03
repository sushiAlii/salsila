package models

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	ID			uint	`gorm:"primaryKey" json:"id"`
	FamilyName	string	`gorm:"not null" json:"familyName"`
	CreatedAt	time.Time		`gorm:"type:timestamptz"`
	UpdatedAt	*time.Time		`gorm:"type:timestamptz"`
}

func CreateFamily(DB *gorm.DB, newFamily *Family) error {
	return DB.Create(newFamily).Error
}

func GetAllFamilies(DB *gorm.DB) ([]Family, error) {
	var familyList []Family

	if err := DB.Find(&familyList).Error; err != nil {
		return nil, err
	}

	return familyList, nil
}

func GetFamilyByID(DB *gorm.DB, id uint) (*Family, error) {
	var family Family

	if err := DB.First(&family, id).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

func UpdateFamilyByID(DB *gorm.DB, id uint, updatedFamily Family) error {
	return DB.Model(&Family{}).Where("id = ?", id).Updates(updatedFamily).Error
}

func DeleteFamilyByID(DB *gorm.DB, id uint) error {
	return DB.Delete(&Family{}, id).Error
}