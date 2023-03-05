package controllers

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     domains.AuthService
	userService domains.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service domains.AuthService,
	userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	user, _ := jwt.userService.GetOneUser(uint(1))
	token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
