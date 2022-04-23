package models

import (
	"time"
)

// User defines a system user
type User struct {
	ID           string
	Disabled     bool
	Username     string
	PasswordHash string
	Name         string
	Created      time.Time
	Updated      time.Time
}
