package models

type PersonsFamily struct {
	ID			uint	`gorm:"primaryKey"`
	FamilyID	uint	`gorm:"not null" json:"familyId"`
	PersonID 	uint	`gorm:"not null" json:"personId"`
	FamilyRole 	string	`json:"familyRole"`
}