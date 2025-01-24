package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/sufimalek/jwtapi/internal/models"
	"github.com/sufimalek/jwtapi/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) (int64, error) {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.UserRepo.DeleteUser(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.UserRepo.ListUsers()
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	// Trim the password
	user.Password = strings.TrimSpace(user.Password)

	// Check if the username already exists
	existingUser, err := s.UserRepo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Hash the password
	log.Printf("Plaintext password during registration: %s", user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	log.Printf("Hashed password during registration: %s", string(hashedPassword))

	// Replace the plaintext password with the hashed password
	user.Password = string(hashedPassword)

	// Create the user
	id, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Set the ID of the newly created user
	user.ID = int(id)

	return user, nil
}
