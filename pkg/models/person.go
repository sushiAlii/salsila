package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	UID			string 		`gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName	string		`gorm:"not null"`
	MiddleName	string		`gorm:"not null"`
	LastName 	string		`gorm:"not null"`
	Gender 		string		`gorm:"not null"`
	Birthday 	time.Time	`gorm:"not null"`
	CreatedAt	time.Time	`gorm:"type:timestamptz"`
	UpdatedAt	*time.Time	`gorm:"type:timestamptz"`
}

func CreatePerson(DB *gorm.DB, person *Person) error {
	return DB.Create(person).Error
}

func GetAllPersons(DB *gorm.DB) ([]Person, error) {
	var personsList []Person

	if err := DB.Find(&personsList).Error; err != nil {
		return nil, err
	}

	return personsList, nil
}

func GetPersonByUID(DB *gorm.DB, uid string) (*Person, error) {
	var person Person

	if err := DB.First(&person, "uid = ?", uid).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func UpdatePersonByUID(DB *gorm.DB, person Person, uid string) error {
	tx := DB.Begin()

	if err := tx.Model(&Person{}).Where("uid = ?", uid).Updates(&person).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func DeletePersonByUID(DB *gorm.DB, uid string) error {
	tx := DB.Begin()

	if err := tx.Delete(&Person{}, "uid = ?", uid).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}