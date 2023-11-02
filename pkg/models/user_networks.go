package models


type UserNetwork struct {
	ID				uint	`gorm:"primaryKey"`
	UserUID			string	`gorm:"not null"`
	SocialNetworkID	uint	`gorm:"not null"`
}
