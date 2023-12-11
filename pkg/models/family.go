package models

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	FamilyName	string		`gorm:"not null" json:"familyName"`
	CreatedAt	time.Time	`gorm:"type:timestamptz"`
	UpdatedAt	*time.Time	`gorm:"type:timestamptz"`
}

type FamilyService interface {
	CreateFamily(*Family) error
	GetAllFamilies() ([]Family, error)
	GetFamilyByID(uint) (*Family, error)
	UpdateFamilyByID(uint, Family) error
	DeleteFamilyByID(uint) error
}

type familyService struct {
	DB *gorm.DB
}

func NewFamilyController(db *gorm.DB) FamilyService {
	return &familyService{DB: db}
}

func (fs *familyService) CreateFamily(newFamily *Family) error {
	return fs.DB.Create(newFamily).Error
}

func (fs *familyService) GetAllFamilies() ([]Family, error) {
	var familyList []Family

	if err := fs.DB.Find(&familyList).Error; err != nil {
		return nil, err
	}

	return familyList, nil
}

func (fs *familyService) GetFamilyByID(id uint) (*Family, error) {
	var family Family

	if err := fs.DB.First(&family, id).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

func (fs *familyService) UpdateFamilyByID(id uint, updatedFamily Family) error {
	return fs.DB.Model(&Family{}).Where("id = ?", id).Updates(updatedFamily).Error
}

func (fs *familyService) DeleteFamilyByID(id uint) error {
	return fs.DB.Delete(&Family{}, id).Error
}