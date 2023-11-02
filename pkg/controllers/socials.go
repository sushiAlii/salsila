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

func CreateSocialNetwork(w http.ResponseWriter, r *http.Request) {
	var newSocialNetwork models.SocialNetwork

	err := json.NewDecoder(r.Body).Decode(&newSocialNetwork)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if newSocialNetwork.Name == "" {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}

	if len(newSocialNetwork.Name) <= 3 {
		http.Error(w, "Name should have more than 3 characters", http.StatusBadRequest)
		return
	}

	if err := models.CreateSocialNetwork(db.DB, &newSocialNetwork); err != nil {
		http.Error(w, "Server Error: Failed to create a new social network", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfuly created a new social network record!",
	})

}

func GetAllSocialNetworks(w http.ResponseWriter, r *http.Request) {
	socialNetworks, err := models.GetAllSocialNetworks(db.DB)

	if err != nil {
		http.Error(w, "Failed to fetch social networks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.SocialNetwork{
		"data": socialNetworks,
	})
}

func GetSocialNetworkByID(w http.ResponseWriter, r *http.Request) {
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
	socialNetwork, err := models.GetSocialNetworkByID(db.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("Social network with ID %d does not exist", id), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch a social network record", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.SocialNetwork {
		"data": socialNetwork,
	}) 	
}

func UpdateSocialNetworkByID(w http.ResponseWriter, r *http.Request) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID format is invalid", http.StatusBadRequest)
		return
	}

	var socialNetworkBody models.SocialNetwork
	if err := json.NewDecoder(r.Body).Decode(&socialNetworkBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	id := uint(parsedId)
	if err := models.UpdateSocialNetworkByID(db.DB, id, socialNetworkBody); err != nil {
		http.Error(w, "Failed to update a social network record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully updated a record",
	})
}

func DeleteSocialNetworkByID(w http.ResponseWriter, r *http.Request) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID format is invalid", http.StatusBadRequest)
		return
	}

	id := uint(parsedId)
	if err := models.DeleteSocialNetworkByID(db.DB, id); err != nil {
		http.Error(w, "Failed to delete a social network record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}