package middlewares

import (
	"github.com/dipeshdulal/clean-gin/lib"
	cors "github.com/rs/cors/wrapper/gin"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler lib.RequestHandler
	logger  lib.Logger
	env     lib.Env
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(handler lib.RequestHandler, logger lib.Logger, env lib.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	debug := m.env.Environment == "development"
	m.handler.Gin.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            debug,
	}))
}
