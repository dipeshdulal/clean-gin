package lib

import "os"

// Env has environment stored
type Env struct {
	ServerPort string
}

// NewEnv creates a new environment
func NewEnv() Env {
	return Env{}
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
}
