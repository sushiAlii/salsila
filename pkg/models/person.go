package models

import (
	"time"
)

type Person struct {
	UID			string 		`gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName	string		`gorm:"not null"`
	MiddleName	string		`gorm:"not null"`
	LastName 	string		`gorm:"not null"`
	Gender 		string		`gorm:"not null"`
	Birthday 	time.Time	`gorm:"not null"`
}