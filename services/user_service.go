package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserService service layer
type UserService struct {
	logger lib.Logger
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger) UserService {
	return UserService{
		logger: logger,
	}
}

// CreateUser call to create the user
func (s UserService) CreateUser() {
	s.logger.Zap.Info("Create user service called.")
}
