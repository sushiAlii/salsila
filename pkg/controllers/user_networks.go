package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

func CreateUserNetwork(w http.ResponseWriter, r *http.Request) {
	var newUserNetwork models.UserNetwork

	if err := json.NewDecoder(r.Body).Decode(&newUserNetwork); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidateCreateUserNetwork(db.DB, &newUserNetwork); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := models.CreateUserNetwork(db.DB, &newUserNetwork); err != nil {
		http.Error(w, "Failed to create a user network", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully linked user and social network",
	})
}

func GetAllUserNetworks(w http.ResponseWriter, r *http.Request) {
	var (
		userNetworks []models.UserNetwork
		err error
	)

	userUid := r.URL.Query().Get("userUid")

	if len(userUid) > 0 {
		userNetworks, err = models.GetUserNetworksByUserUID(db.DB, userUid)
	} else {
		userNetworks, err = models.GetAllUserNetworks(db.DB)
	}

	if err != nil {
		http.Error(w, "Failed to fetch user network records", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.UserNetwork{
		"data": userNetworks,
	})
}

func UpdateUserNetworkByID(w http.ResponseWriter, r *http.Request) {
	var userNetwork models.UserNetwork

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

	if err := json.NewDecoder(r.Body).Decode(&userNetwork); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidateUpdateUserNetwork(db.DB, &userNetwork); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id := uint(parsedId)
	if err := models.UpdateUserNetworkByID(db.DB, &userNetwork, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update a user network record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully updated a user network!",
	})
}

func DeleteUserNetworkByID(w http.ResponseWriter, r *http.Request) {
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
	if err := models.DeleteUserNetworkByID(db.DB, id); err != nil {
		http.Error(w, "Failed to delete a user network record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}