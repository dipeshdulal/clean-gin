package lib

import "os"

// Env has environment stored
type Env struct {
	ServerPort  string
	Environment string
	LogOutput   string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	JWTSecret   string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.Environment = os.Getenv("ENV")
	env.LogOutput = os.Getenv("LOG_OUTPUT")

	env.DBUsername = os.Getenv("DB_USER")
	env.DBPassword = os.Getenv("DB_PASS")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")

	env.JWTSecret = os.Getenv("JWT_SECRET")
}
