package lib

import "os"

// Env has environment stored
type Env struct {
	ServerPort  string
	Environment string
	LogOutput   string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
	env.Environment = os.Getenv("Environment")
	env.LogOutput = os.Getenv("LogOutput")
}
