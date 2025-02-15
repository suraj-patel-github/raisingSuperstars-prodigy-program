package service

import (
	"context"
	"prodigy-program/repos"
)

// Define the UserService Interface
type UserService interface {
	RegisterUser(ctx context.Context, name, email string) (int, error)
}

// Implement the UserService Struct
type userService struct {
	repo repos.UserRepository
}

// Constructor for UserService
func NewUserService(repo repos.UserRepository) UserService {
	return &userService{repo: repo}
}

// Business Logic for Registering a User
func (s *userService) RegisterUser(ctx context.Context, name, email string) (int, error) {
	return s.repo.CreateUser(ctx, name, email)
}