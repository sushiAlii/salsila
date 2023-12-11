package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sushiAlii/salsila/pkg/models"
)

type AuthController struct {
	authService models.AuthService
	userService models.UserService
}

func NewAuthController(authService models.AuthService, userService models.UserService) *AuthController {
	return &AuthController{authService: authService, userService: userService}
}

func (ac *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ac.authService.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := ac.authService.CreateToken(user.UID)
	if err != nil {
		http.Error(w, "Failed to create a token", http.StatusInternalServerError)
		return
	}

	if err := ac.authService.SaveAuth(user.UID, token); err != nil {
		http.Error(w, "Failed to save authentication token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

func (ac *AuthController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Refresh-Token")
	if(refreshToken == "") {
		http.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}

	if err := ac.authService.LogoutUser(refreshToken); err != nil {
		http.Error(w, "Could not logout", http.StatusInternalServerError)
		return
	}
	
	w.Header().Del("Refresh-Token")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successfully",
	})
}

func (ac *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := ac.userService.ValidateUser(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}


	if err := ac.authService.RegisterUser(&newUser); err != nil {
		http.Error(w, "Failed to register a user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully created a User!",
	})
}

func (ac *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Refresh-Token")

	if refreshToken == "" {
		http.Error(w, "Refresh token does not exist", http.StatusBadRequest)
		return
	}

	newToken, err := ac.authService.Refresh(refreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": newToken.AccessToken,
		"refresh_token": newToken.RefreshToken,
	})
}