package repos

import (
	"context"
	"database/sql"
)

// Define the UserRepository Interface
type UserRepository interface {
	CreateUser(ctx context.Context, name, email string) (int, error)
}

// Implement the Repository Struct
type userRepository struct {
	db *sql.DB
}

// Constructor for UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Store a New User in the Database
func (r *userRepository) CreateUser(ctx context.Context, name, email string) (int, error) {
	var userID int
	query := `INSERT INTO Users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, name, email).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
