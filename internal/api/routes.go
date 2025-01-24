package api

import (
	"database/sql"

	"github.com/sufimalek/jwtapi/internal/api/handlers"
	"github.com/sufimalek/jwtapi/internal/api/middleware"
	"github.com/sufimalek/jwtapi/internal/repository"
	"github.com/sufimalek/jwtapi/internal/service"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	router.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTMiddleware)
	protected.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	protected.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	protected.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	return router
}
