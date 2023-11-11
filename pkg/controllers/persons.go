package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		fmt.Printf("ERROR!! : %v", err.Error())

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidatePerson(db.DB, &person); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := models.CreatePerson(db.DB, &person); err != nil {
		http.Error(w, "Failed to create a new person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a new person!",
	})
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	personsList, err := models.GetAllPersons(db.DB)
	if err != nil {
		http.Error(w, "Failed to fetch persons list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.Person{
		"data": personsList,
	})
}

func GetPersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	person, err := models.GetPersonByUID(db.DB, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to retrieve person data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.Person{
		"data": person,
	})
}

func UpdatePersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	var updatedPerson models.Person
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := models.ValidatePerson(db.DB, &updatedPerson); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := models.UpdatePersonByUID(db.DB, updatedPerson, uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update Person record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully updated a person",
	})
}

func DeletePersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	if err := models.DeletePersonByUID(db.DB, uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}