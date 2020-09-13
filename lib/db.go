package lib

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Database modal
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, logger Logger) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open("mysql", url)

	if err != nil {
		logger.Zap.Info("Url: ", url)
		logger.Zap.Panic(err)
	}

	logger.Zap.Info("Database connection established")

	return Database{
		DB: db,
	}
}
