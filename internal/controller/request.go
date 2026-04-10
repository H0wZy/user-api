package controller

import (
	"time"

	"github.com/H0wZy/user-api/internal/types"
	"github.com/H0wZy/user-api/internal/utils"
)

type CreateUserRequest struct {
	Username  string         `json:"username" binding:"required"`
	Email     string         `json:"email" binding:"required"`
	FirstName string         `json:"first_name" binding:"required"`
	LastName  string         `json:"last_name" binding:"required"`
	Password  string         `json:"password" binding:"required,min=8"`
	Phone     string         `json:"phone" binding:"required"`
	Gender    []types.Gender `json:"gender" binding:"required"`
	BirthDate *time.Time     `json:"birth_date,omitempty"`

	// DONT NEED:
	// LoginMethod  string         `json:"login_method,omitempty"`
	// LastLoginAt  *time.Time     `json:"last_login_at,omitempty"`
	// LastLogoutAt *time.Time     `json:"last_logout_at,omitempty"`
	// IsUserActive bool           `json:"is_user_active,omitempty"`
}

type UpdateUserRequest struct {
	Username  string         `json:"username" binding:"required"`
	Email     string         `json:"email" binding:"required"`
	FirstName string         `json:"first_name" binding:"required"`
	LastName  string         `json:"last_name" binding:"required"`
	Password  string         `json:"password" binding:"required,min=8"`
	Phone     string         `json:"phone" binding:"required"`
	Gender    []types.Gender `json:"gender" binding:"required"`
}

func (r *UpdateUserRequest) ValidateGender() error {
	return utils.ValidateGender(r.Gender)
}

func (r *CreateUserRequest) ValidateGender() error {
	return utils.ValidateGender(r.Gender)
}
