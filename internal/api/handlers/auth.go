package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sufimalek/jwtapi/internal/service"
	"github.com/sufimalek/jwtapi/internal/utils"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	utils.Log.Info("AuthHandler called")

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for user: %s", creds.Username)

	token, err := h.AuthService.Authenticate(creds.Username, creds.Password)
	if err != nil {
		log.Printf("Authentication failed for user: %s, error: %v", creds.Username, err)
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	log.Printf("Login successful for user: %s", creds.Username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
