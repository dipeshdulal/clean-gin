package controllers

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	Get  func(c *gin.Context)
	Post func(c *gin.Context)
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		Get:  getUserController(userService, logger),
		Post: postUserController(),
	}
}

func getUserController(userService services.UserService, logger lib.Logger) func(*gin.Context) {
	return func(c *gin.Context) {
		userService.CreateUser()
		logger.Zap.Info("Get user route called")
		c.JSON(200, "Get User")
	}
}

func postUserController() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Post User")
	}
}
