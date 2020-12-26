package models

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name"`
	Email        *string        `json:"email"`
	Age          uint8          `json:"age"`
	Birthday     *time.Time     `json:"time"`
	MemberNumber sql.NullString `json:"member_number"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// TableName gives table name of model
func (u User) TableName() string {
	return "users"
}
