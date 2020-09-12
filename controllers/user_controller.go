package controllers

import "github.com/gin-gonic/gin"

// UserController data type
type UserController struct {
	Get  func(c *gin.Context)
	Post func(c *gin.Context)
}

// NewUserController creates new user controller
func NewUserController() UserController {
	return UserController{
		Get:  getUserController(),
		Post: postUserController(),
	}
}

func getUserController() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Get User")
	}
}

func postUserController() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Post User")
	}
}
