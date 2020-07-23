package entities

import (
	"mcs_service/models/auxiliary"
	"time"
)

type User struct {
	auxiliary.BaseModel
	Username   string     `json:"username"            gorm:"not null"`
	Password   string     `json:"-"                   gorm:"not null"`
	LastLogin  *time.Time `json:"last_login"`
	UserRoleID uint       `json:"user_role_id"`
	UserRole   *UserRole  `json:"user_role,omitempty"`
}
