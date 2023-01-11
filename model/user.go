package model

import "time"

type User struct {
	UserID    string
	Username  string
	Password  string
	Email     string
	Verified  bool
	CreatedAt time.Time
}
