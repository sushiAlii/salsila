package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

func CreatePersonFamily(w http.ResponseWriter, r *http.Request) {
	var personFamily models.PersonsFamily

	if err := json.NewDecoder(r.Body).Decode(&personFamily); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.CreatePersonsFamily(db.DB, &personFamily); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a person family",
	})
}

func GetAllPersonsFamilies(w http.ResponseWriter, r *http.Request) {
	var (
		parsedFamilyId int
		personFamilies []models.PersonsFamily
		err error
	)

	familyIdParam := r.URL.Query().Get("familyId")
	personUidParam := r.URL.Query().Get("personUid")

	parsedFamilyId, err = strconv.Atoi(familyIdParam)
	if err != nil {
		http.Error(w, "Invalid family ID", http.StatusBadRequest)
		return
	}

	familyId := uint(parsedFamilyId)

	if familyId != 0 {
		personFamilies, err = models.GetPersonsFamiliesByFamilyID(db.DB, familyId)
	} else if len(personUidParam) > 0 {
		personFamilies, err = models.GetPersonsFamiliesByPersonUID(db.DB, personUidParam)
	} else {
		personFamilies, err = models.GetAllPersonsFamilies(db.DB)
	}

	if err != nil {
		http.Error(w, "Failed to fetch person families list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.PersonsFamily{
		"data": personFamilies,
	})
}

func GetPersonFamilyByID(w http.ResponseWriter, r *http.Request) {
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
	personFamily, err := models.GetPersonFamilyByID(db.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch person family record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.PersonsFamily {
		"data": personFamily,
	})
}

func UpdatePersonFamilyByID(w http.ResponseWriter, r *http.Request) {
	var personFamily models.PersonsFamily

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

	if err := json.NewDecoder(r.Body).Decode(&personFamily); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := uint(parsedId)
	if err := models.UpdatePersonFamilyByID(db.DB, &personFamily, id); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err.Error()), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully updated a person family record",
	})
}

func DeletePersonsFamilyByID(w http.ResponseWriter, r *http.Request) {
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
	if err := models.DeletePersonsFamilyByID(db.DB, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete person family record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}