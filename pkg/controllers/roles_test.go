package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/sushiAlii/salsila/pkg/models"
	"github.com/sushiAlii/salsila/test/mocks"
	"gorm.io/gorm"
)

func TestCreateRole(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	mockRole := &models.Role{Name: "Janitor", Description: "Cleans the office and makes me coffee",}
	mockRoleService.On("CreateRole", mockRole).Return(nil)

	body, _ := json.Marshal(mockRole)
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(roleController.CreateRole)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockRoleService.AssertExpectations(t)

}

func TestGetAllRoles(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	mockRolesList := []models.Role{
		{ID: 1, Name: "Mocked Admin", Description: "Mocked Admin role"},
		{ID: 2, Name: "Mocked Dev", Description: "Mocked Dev role"},
		{ID: 3, Name: "Mocked User", Description: "Mocked User role"},
	}

	mockResponse := map[string][]models.Role{
		"data": mockRolesList,
	}

	mockRoleService.On("GetAllRoles").Return(mockRolesList, nil)

	req, _ := http.NewRequest("GET", "/roles", nil)
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(roleController.GetAllRoles)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string][]models.Role
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, response)

	mockRoleService.AssertExpectations(t)
}

func TestGetRoleByID(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)
	
	roleID := uint(1)
	mockRole := &models.Role{ID: int(roleID), Name: "Janitor", Description: "Cleans the office and makes me coffee",}
	mockResponse := map[string]*models.Role{
		"data": mockRole,
	}

	mockRoleService.On("GetRoleByID", roleID).Return(mockRole, nil)

	req, _ := http.NewRequest("GET", "/roles/1", nil)
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/roles/{id:[0-9]+}", roleController.GetRoleByID).Methods("GET")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]*models.Role
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.Equal(t, mockResponse, response)

	mockRoleService.AssertExpectations(t)
}

func TestUpdateRoleByID_Success(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	roleID := uint(1)
	updatedRole := &models.Role{Name: "Senior Janitor", Description: "Oversees cleaning"}

	mockRoleService.On("UpdateRoleByID", roleID, *updatedRole).Return(nil)

	body, _ := json.Marshal(updatedRole)
	req, _ := http.NewRequest("PATCH", "/roles/1", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/roles/{id:[0-9]+}", roleController.UpdateRoleByID).Methods("PATCH")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	mockRoleService.AssertExpectations(t)
}

func TestUpdateRoleByID_Failure(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	roleID := uint(999)
	updatedRole := &models.Role{Name: "Invalid", Description: "Invalid Description"}

	mockRoleService.On("UpdateRoleByID", roleID, *updatedRole).Return(errors.New("Role ID not found"))

	body, _ := json.Marshal(updatedRole)
	req, _ := http.NewRequest("PATCH", "/roles/999", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/roles/{id:[0-9]+}", roleController.UpdateRoleByID).Methods("PATCH")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	mockRoleService.AssertExpectations(t)
}

func TestDeleteRoleByID_Success(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	roleID := uint(1)
	mockRoleService.On("DeleteRoleByID", roleID).Return(nil)

	req, _ := http.NewRequest("DELETE", "/roles/1", nil)
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/roles/{id:[0-9]+}", roleController.DeleteRoleByID).Methods("DELETE")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNoContent, recorder.Code)
	mockRoleService.AssertExpectations(t)
}

func TestDeleteRoleByID_Failure(t *testing.T) {
	mockRoleService := new(mocks.MockRoleService)
	roleController := NewRoleController(mockRoleService)

	roleID := uint(999)
	mockRoleService.On("DeleteRoleByID", roleID).Return(gorm.ErrRecordNotFound)

	req, _ := http.NewRequest("DELETE", "/roles/999", nil)
	recorder := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/roles/{id:[0-9]+}", roleController.DeleteRoleByID).Methods("DELETE")
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	mockRoleService.AssertExpectations(t)
}