package controllers

import (
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	Get  func(c *gin.Context)
	Post func(c *gin.Context)
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService) UserController {
	return UserController{
		Get:  getUserController(userService),
		Post: postUserController(),
	}
}

func getUserController(userService services.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		userService.CreateUser()
		c.JSON(200, "Get User")
	}
}

func postUserController() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Post User")
	}
}
