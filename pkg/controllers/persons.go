package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

type PersonController struct {
	service models.PersonService
}

func NewPersonController(service models.PersonService) *PersonController {
	return &PersonController{service: service}
}

func (pc *PersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		fmt.Printf("ERROR!! : %v", err.Error())

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidatePerson(&person); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := pc.service.CreatePerson(&person); err != nil {
		http.Error(w, "Failed to create a new person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a new person!",
	})
}

func (pc *PersonController) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	personsList, err := pc.service.GetAllPersons()
	if err != nil {
		http.Error(w, "Failed to fetch persons list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.Person{
		"data": personsList,
	})
}

func (pc *PersonController) GetPersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	person, err := pc.service.GetPersonByUID(uid)
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

func (pc *PersonController) UpdatePersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	var updatedPerson models.Person
	if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if err := models.ValidatePerson(&updatedPerson); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := pc.service.UpdatePersonByUID(updatedPerson, uid); err != nil {
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

func (pc *PersonController) DeletePersonByUID(w http.ResponseWriter, r *http.Request) {
	uid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "UID not provided", http.StatusBadRequest)
		return
	}

	if err := pc.service.DeletePersonByUID(uid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}