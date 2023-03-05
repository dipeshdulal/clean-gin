package domains

import "github.com/dipeshdulal/clean-gin/models"

type AuthService interface {
	Authorize(tokenString string) (bool, error)
	CreateToken(models.User) string
}
