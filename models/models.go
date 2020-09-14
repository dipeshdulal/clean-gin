package models

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"go.uber.org/fx"
)

// Module exported from models package
var Module = fx.Options(
	fx.Provide(NewMigrations),
)

// Migrations migration data type
type Migrations struct {
	db     lib.Database
	logger lib.Logger
}

// NewMigrations creates new migrations instance
func NewMigrations(db lib.Database, logger lib.Logger) Migrations {
	return Migrations{
		db:     db,
		logger: logger,
	}
}

// Migrate function to call when migrating
func (m Migrations) Migrate() {
	m.logger.Zap.Info("Automigrating schemas.")
	m.db.DB.AutoMigrate(&User{})
}
