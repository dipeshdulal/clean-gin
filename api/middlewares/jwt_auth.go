package middlewares

import (
	"net/http"
	"strings"

	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middleware for jwt authentication
type JWTAuthMiddleware struct {
	service domains.AuthService
	logger  lib.Logger
}

// NewJWTAuthMiddleware creates new jwt auth middleware
func NewJWTAuthMiddleware(
	logger lib.Logger,
	service domains.AuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		service: service,
		logger:  logger,
	}
}

// Setup sets up jwt auth middleware
func (m JWTAuthMiddleware) Setup() {}

// Handler handles middleware functionality
func (m JWTAuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := m.service.Authorize(authToken)
			if authorized {
				c.Next()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			m.logger.Error(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		c.Abort()
	}
}
