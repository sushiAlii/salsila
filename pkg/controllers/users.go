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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var (
		users	[]models.User
		err		error
	)

	users, err = models.GetAllUsers(db.DB)
	
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

func AttachPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO??")

	var body struct {
		PersonUid	string
	}

	userUid, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "User's UID not provided", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := models.AttachPerson(db.DB, body.PersonUid, userUid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to attach person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfuly attached a person to the user",
	})
}

func DeleteUserByUID(w http.ResponseWriter, r *http.Request) {
	uidParam, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	if err := models.DeleteUserByUID(db.DB, uidParam); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to Delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}