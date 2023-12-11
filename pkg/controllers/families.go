package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

type FamilyController struct {
	service models.FamilyService
}

func NewFamilyController(service models.FamilyService) *FamilyController {
	return &FamilyController{service: service}
}

func (fc *FamilyController) CreateFamily(w http.ResponseWriter, r *http.Request) {
	var newFamily models.Family

	fmt.Printf("Response Body %v \n", r.Body)

	if err := json.NewDecoder(r.Body).Decode(&newFamily); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)

		fmt.Printf("Error occured! %v \n", newFamily)
		return
	}

	if newFamily.FamilyName == "" {
		http.Error(w, "Family name is required", http.StatusUnprocessableEntity)
		return
	}

	if len(newFamily.FamilyName) <= 3 {
		http.Error(w, "Family name should be more than 3 characters", http.StatusUnprocessableEntity)
		return
	}

	if err := fc.service.CreateFamily(&newFamily); err != nil {
		http.Error(w, "Failed to create a new family record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a family",
	})
}

func (fc *FamilyController) GetAllFamilies(w http.ResponseWriter, r *http.Request) {
	familiesList, err := fc.service.GetAllFamilies()

	if err != nil {
		http.Error(w, "Failed to fetch families", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.Family {
		"data": familiesList,
	})
}

func (fc *FamilyController) GetFamilyByID(w http.ResponseWriter, r *http.Request) {
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
	familyData, err := fc.service.GetFamilyByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Family record with ID %d is not found", id), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch a family record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.Family{
		"data": familyData,
	})
}

func (fc *FamilyController) UpdateFamilyByID(w http.ResponseWriter, r *http.Request) {
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

	var updatedFamily models.Family

	if err := json.NewDecoder(r.Body).Decode(&updatedFamily); err != nil {
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := uint(parsedId)
	if err := fc.service.UpdateFamilyByID(id, updatedFamily); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Family with ID %d does not exist", id), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update family record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string {
		"message": "Family updated successfully!",
	})
}

func (fc *FamilyController) DeleteFamilyByID(w http.ResponseWriter, r *http.Request) {
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
	if err := fc.service.DeleteFamilyByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Family ID %d does not exist", id), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete a family record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}