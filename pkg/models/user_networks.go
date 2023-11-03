package models

import "time"


type UserNetwork struct {
	ID				uint		`gorm:"primaryKey"`
	UserUID			string		`gorm:"not null"`
	SocialNetworkID	uint		`gorm:"not null"`
	CreatedAt		time.Time	`gorm:"type:timestamptz"`
	UpdatedAt		*time.Time	`gorm:"type:timestamptz"`
}
