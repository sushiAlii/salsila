package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	var (
		users	[]models.User
		err		error
	)

	if email != "" {
		users, err = models.GetUsersByEmail(db.DB, email)
	} else {
		users, err = models.GetAllUsers(db.DB)
	}

	if err != nil {
		http.Error(w, "Failed to fetch list of users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.User{
		"data": users,
	})
}

func GetUserByUID(w http.ResponseWriter, r *http.Request) {
	uidParam, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	userData, err := models.GetUserByUID(db.DB, uidParam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User does not exist", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.User{
		"data": userData,
	})
}