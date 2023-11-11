package models

import "gorm.io/gorm"

type Role struct {
	ID			int		`gorm:"primaryKey" json:"id"`
	Name		string	`gorm:"not null" json:"name"`
	Description	string	`json:"description"`
}

func CreateRole(DB *gorm.DB, role *Role) error {
	return DB.Create(role).Error
}

func GetAllRoles(DB *gorm.DB) ([]Role, error) {
	var rolesList []Role

	if err := DB.Find(&rolesList).Error; err != nil {
		return nil, err
	}
	
	return rolesList, nil
}

func GetRoleByID(DB *gorm.DB, id uint) (*Role, error) {
	var roleData Role

	if err := DB.First(&roleData, id).Error; err != nil {
		return nil, err
	}
	return &roleData, nil
}

func UpdateRoleByID(DB *gorm.DB, id uint, updatedRole Role) error {
	return DB.Model(&Role{}).Where("id = ?", id).Updates(updatedRole).Error
}

func DeleteRoleByID(DB *gorm.DB, id uint) error {
	return DB.Delete(&Role{}, id).Error
}