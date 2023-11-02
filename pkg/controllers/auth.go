package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.ValidateUser(db.DB, &newUser); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}


	if err := models.RegisterUser(db.DB, &newUser); err != nil {
		http.Error(w, "Failed to register a user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a User!",
	})
}