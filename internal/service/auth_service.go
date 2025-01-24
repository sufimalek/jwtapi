package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/sufimalek/jwtapi/internal/repository"
	"github.com/sufimalek/jwtapi/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
	log.Printf("Attempting to authenticate user: %s", username)

	// Trim the password
	password = strings.TrimSpace(password)

	// Find the user by username
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		log.Printf("User not found: %s", username)
		return "", fmt.Errorf("user not found")
	}

	log.Printf("User found: %+v", user)

	// Compare the hashed password with the provided password
	utils.Log.Info("testing logs......")
	log.Printf("Hashed password from database: %s", user.Password)
	log.Printf("Provided password: %s", password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Password mismatch for user: %s", username)
		return "", fmt.Errorf("invalid password")
	}

	log.Printf("Password matched for user: %s", username)

	// Generate a JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		log.Printf("Failed to generate token for user: %s, error: %v", username, err)
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	log.Printf("Token generated successfully for user: %s", username)
	return token, nil
}
