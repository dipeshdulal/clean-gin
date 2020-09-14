package routes

import (
	"github.com/dipeshdulal/clean-gin/controllers"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetUser)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.POST("/user", s.userController.SaveUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(handler lib.RequestHandler, userController controllers.UserController, logger lib.Logger) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}
