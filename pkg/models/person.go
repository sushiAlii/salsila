package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	UID			string 		`gorm:"type:uuid;default:uuid_generate_v4()" json:"uid"`
	FirstName	string		`gorm:"not null" json:"firstName"`
	MiddleName	string		`gorm:"not null" json:"middleName"`
	LastName 	string		`gorm:"not null" json:"lastName"`
	Gender 		string		`gorm:"not null" json:"gender"`
	Birthday 	string		`gorm:"type:date;not null" json:"birthday"`
	CreatedAt	time.Time	`gorm:"type:timestamptz" json:"-"`
}

var GenderEnum = [2]string{"Male", "Female"}

func (Person) TableName() string {
	return "persons"
}

func ValidatePerson(DB *gorm.DB, person *Person) error {
	if person.FirstName == "" {
		return ErrFirstNameRequired
	}

	if person.LastName == "" {
		return ErrLastNameRequired
	}

	if person.MiddleName == "" {
		return ErrMiddleNameRequired
	}

	if person.Gender == "" {
		return ErrGenderRequired
	}

	_, err := time.Parse("2006-01-02", person.Birthday)
	if err != nil {
		return ErrBirthdayInvalid
	}

	isValidGender := false
	for _, validGender := range GenderEnum {
		if person.Gender == validGender{
			isValidGender = true
			break
		}
	}

	if !isValidGender {
		return ErrGenderInvalid
	}

	return nil
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