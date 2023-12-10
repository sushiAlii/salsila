package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

type RoleController struct {
	service models.RoleService
}

func NewRoleController(service models.RoleService) *RoleController {
	return &RoleController{service: service}
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var newRole models.Role

	err := json.NewDecoder(r.Body).Decode(&newRole)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if newRole.Name == "" {
		http.Error(w, "Role name is a required field", http.StatusUnprocessableEntity)

		return
	}

	if len(newRole.Name) <= 3 {
		http.Error(w, "Role name should have more than 3 characters", http.StatusUnprocessableEntity)
		return
	}

	

	if err := rc.service.CreateRole(&newRole); err != nil {
		http.Error(w, "Failed to create a new role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Role created successfully!",
	})
}

func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.service.GetAllRoles()

	if err != nil {
		http.Error(w, "Failed to retrieve roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.Role{
		"data": roles,
	})
}

func (rc *RoleController) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	id := uint(parsedId)
	roleData, err := rc.service.GetRoleByID(id)
	if err != nil {
		http.Error(w, "Role data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.Role{
		"data": roleData,
	})
}

func (rc *RoleController) UpdateRoleByID(w http.ResponseWriter, r *http.Request) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedRole models.Role
	if err := json.NewDecoder(r.Body).Decode(&updatedRole); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := uint(parsedId)
	if err := rc.service.UpdateRoleByID(id, updatedRole); err != nil {
		http.Error(w, "Failed to update role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Role updated successfully!",
	})
}

func (rc *RoleController) DeleteRoleByID(w http.ResponseWriter, r *http.Request) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	id := uint(parsedId)
	if err := rc.service.DeleteRoleByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Role record not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete a role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}