package model

import (
	"time"

	"github.com/H0wZy/user-api/internal/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Email        string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	FirstName    string         `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string         `gorm:"type:varchar(100);not null" json:"last_name"`
	FullName     string         `gorm:"type:varchar(200);not null" json:"full_name"`
	Password     string         `gorm:"not null" json:"-"`
	Phone        string         `gorm:"type:varchar(20);not null" json:"phone"`
	Gender       []types.Gender `gorm:"serializer:json;not null" json:"gender"`
	BirthDate    *time.Time     `json:"birth_date,omitempty"`
	LoginMethod  string         `gorm:"type:varchar(20);" json:"login_method"`
	LastLoginAt  *time.Time     `json:"last_login_at,omitempty"`
	LastLogoutAt *time.Time     `json:"last_logout_at,omitempty"`
	IsUserActive bool           `gorm:"default:true" json:"is_user_active"`
	// CreatedAt    time.Time      `json:"created_at"`
	// UpdatedAt    time.Time      `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserResponse struct {
	User
}
