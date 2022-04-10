package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	lib.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
