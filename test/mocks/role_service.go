package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/sushiAlii/salsila/pkg/models"
)

type MockRoleService struct {
	mock.Mock
}

func (m *MockRoleService) CreateRole(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockRoleService) GetAllRoles() ([]models.Role, error) {
	args := m.Called()
	return args.Get(0).([]models.Role), args.Error(1)
}

func (m *MockRoleService) GetRoleByID(id uint) (*models.Role, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *MockRoleService) UpdateRoleByID(id uint, updatedRole models.Role) error {
	args := m.Called(id, updatedRole)
	return args.Error(0)
}

func (m *MockRoleService) DeleteRoleByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}