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

type PersonService interface {
	CreatePerson(*Person) error
	GetAllPersons() ([]Person, error)
	GetPersonByUID(string) (*Person, error)
	UpdatePersonByUID(Person, string) error
	DeletePersonByUID(string) error
}

type personService struct {
	DB *gorm.DB
}

func NewPersonService(db *gorm.DB) PersonService {
	return &personService{DB: db}
}

func ValidatePerson(person *Person) error {
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

func (ps *personService) CreatePerson(person *Person) error {
	return ps.DB.Create(person).Error
}

func (ps *personService) GetAllPersons() ([]Person, error) {
	var personsList []Person

	if err := ps.DB.Find(&personsList).Error; err != nil {
		return nil, err
	}

	return personsList, nil
}

func (ps *personService) GetPersonByUID(uid string) (*Person, error) {
	var person Person

	if err := ps.DB.First(&person, "uid = ?", uid).Error; err != nil {
		return nil, err
	}

	return &person, nil
}

func (ps *personService) UpdatePersonByUID(person Person, uid string) error {
	tx := ps.DB.Begin()

	if err := tx.Model(&Person{}).Where("uid = ?", uid).Updates(&person).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func (ps *personService) DeletePersonByUID(uid string) error {
	tx := ps.DB.Begin()

	if err := tx.Delete(&Person{}, "uid = ?", uid).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}