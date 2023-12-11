package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/models"
	"gorm.io/gorm"
)

type UserController struct {
	service models.UserService
}

func NewUserController(service models.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var (
		users	[]models.User
		err		error
	)

	users, err = uc.service.GetAllUsers()
	
	if err != nil {
		http.Error(w, "Failed to fetch list of users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.User{
		"data": users,
	})
}

func (uc *UserController) GetUserByUID(w http.ResponseWriter, r *http.Request) {
	uidParam, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	userData, err := uc.service.GetUserByUID(uidParam)
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

func (uc *UserController) AttachPerson(w http.ResponseWriter, r *http.Request) {
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

	if err := uc.service.AttachPerson(body.PersonUid, userUid); err != nil {
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

func (uc *UserController) DeleteUserByUID(w http.ResponseWriter, r *http.Request) {
	uidParam, ok := mux.Vars(r)["uid"]
	if !ok {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	if err := uc.service.DeleteUserByUID(uidParam); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to Delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}