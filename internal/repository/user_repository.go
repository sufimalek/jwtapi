package repository

import (
	"database/sql"

	"github.com/sufimalek/jwtapi/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, password, email, created_at, updated_at FROM users WHERE username = ?`
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (r *UserRepository) CreateUser(user *models.User) error {
// 	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
// 	_, err := r.DB.Exec(query, user.Username, user.Password, user.Email)
// 	return err
// }

func (r *UserRepository) CreateUser(user *models.User) (int64, error) {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	result, err := r.DB.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted user
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = ?, email = ? WHERE id = ?`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, username, email, created_at, updated_at FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
