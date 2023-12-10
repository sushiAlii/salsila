package models

import "gorm.io/gorm"

type Role struct {
	ID			int		`gorm:"primaryKey" json:"id"`
	Name		string	`gorm:"not null" json:"name"`
	Description	string	`json:"description"`
}

type RoleService interface {
	CreateRole(*Role) error
	GetAllRoles() ([]Role, error)
	GetRoleByID(uint) (*Role, error)
	UpdateRoleByID(uint, Role) error
	DeleteRoleByID(uint) error
}

type roleService struct {
	DB *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{DB: db}
}

func (rs *roleService) CreateRole(role *Role) error {
	return rs.DB.Create(role).Error	
}

func (rs *roleService) GetAllRoles() ([]Role, error) {
	var rolesList []Role

	if err := rs.DB.Find(&rolesList).Error; err != nil {
		return nil, err
	}

	return rolesList, nil
}

func (rs *roleService) GetRoleByID(id uint) (*Role, error) {
	var roleData Role

	if err := rs.DB.First(&roleData, id).Error; err != nil {
		return nil, err
	}

	return &roleData, nil
}

func (rs *roleService) UpdateRoleByID(id uint, updatedRole Role) error {
	return rs.DB.Model(&Role{}).Where("id = ?", id).Updates(updatedRole).Error
}

func (rs *roleService) DeleteRoleByID(id uint) error {
	return rs.DB.Delete(&Role{}, id).Error
}