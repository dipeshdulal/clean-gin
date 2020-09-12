package services

import "fmt"

// UserService service layer
type UserService struct {
}

// NewUserService creates a new userservice
func NewUserService() UserService {
	return UserService{}
}

// CreateUser call to create the user
func (s UserService) CreateUser() {
	fmt.Println("Create user service called.")
}
