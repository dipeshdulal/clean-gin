package routers

import (
	"fmt"

	"github.com/dipeshdulal/clean-gin/controllers"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type UserRoutes struct {
	handler        lib.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	fmt.Println("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.Get)
		api.POST("/user", s.userController.Post)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(handler lib.RequestHandler, userController controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
