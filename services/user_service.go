package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id uint) (user models.User, err error) {
	return user, s.repository.Find(&user, id).Error
}

// GetAllUser get all the user
func (s UserService) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// CreateUser call to create the user
func (s UserService) CreateUser(user models.User) error {
	return s.repository.Create(&user).Error
}

// UpdateUser updates the user
func (s UserService) UpdateUser(user models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}
