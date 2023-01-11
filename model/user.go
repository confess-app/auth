package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `user`
func (User) TableName() string {
	return "user"
}
