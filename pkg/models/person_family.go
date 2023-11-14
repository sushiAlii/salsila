package models

import (
	"time"

	"gorm.io/gorm"
)

type PersonsFamily struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	FamilyID	uint		`gorm:"not null" json:"familyId"`
	PersonUID 	string		`gorm:"not null" json:"personId"`
	FamilyRole 	string		`gorm:"not null" json:"familyRole"`
	CreatedAt	time.Time	`gorm:"type:timestamptz" json:"-"`
	UpdatedAt	*time.Time	`gorm:"type:timestamptz" json:"-"`
}

func validateCreatePersonsFamily(personsFamily *PersonsFamily) error {
	if personsFamily.FamilyID == 0 {
		return ErrFamilyIDRequired
	}

	if personsFamily.PersonUID == "" {
		return ErrPersonUIDRequired
	}

	if personsFamily.FamilyRole == "" {
		return ErrFamilyRoleRequired
	}

	return nil
}

func validateUpdatePersonsFamily(personsFamily *PersonsFamily) error {
	if personsFamily.FamilyRole == "" {
		return ErrFamilyRoleRequired
	}

	return nil
}

func CreatePersonsFamily(DB *gorm.DB, personsFamily *PersonsFamily) error {
	tx := DB.Begin()

	if err := validateCreatePersonsFamily(personsFamily); err != nil {
		tx.Rollback()
		
		return err
	}

	if err := tx.Create(personsFamily).Error; err != nil{
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func GetAllPersonsFamilies(DB *gorm.DB) ([]PersonsFamily, error) {
	var personsFamilies []PersonsFamily

	if err := DB.Find(&personsFamilies).Error; err != nil {
		return nil, err
	}

	return personsFamilies, nil
}

func GetPersonsFamiliesByFamilyID(DB *gorm.DB, familyId uint) ([]PersonsFamily, error) {
	var personsFamilies []PersonsFamily

	if err := DB.Find(&personsFamilies, "family_id = ?", familyId).Error; err != nil {
		return nil, err
	}

	return personsFamilies, nil
}

func GetPersonsFamiliesByPersonUID(DB *gorm.DB, personUid string) ([]PersonsFamily, error) {
	var personsFamilies []PersonsFamily

	if err := DB.Find(&personsFamilies, "person_uid = ?", personUid).Error; err != nil {
		return nil, err
	}

	return personsFamilies, nil
}

func GetPersonFamilyByID(DB *gorm.DB, id uint) (*PersonsFamily, error){
	var personsFamily PersonsFamily

	if err := DB.Find(&personsFamily, id).Error; err != nil {
		return nil, err
	}

	return &personsFamily, nil
}

func UpdatePersonFamilyByID(DB *gorm.DB, personsFamily *PersonsFamily, id uint) error {
	if err := validateUpdatePersonsFamily(personsFamily); err != nil {
		return err
	}

	tx := DB.Begin()

	if err := tx.Model(&PersonsFamily{}).Updates(&personsFamily).Where("id = ?", id).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func DeletePersonsFamilyByID(DB *gorm.DB, id uint) error {
	tx := DB.Begin()

	if err := tx.Delete(&PersonsFamily{}, id).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}