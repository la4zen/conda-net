package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Login     string    `json:"login,omitempty"`
	Password  string    `json:"password,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty" gorm:"type:time"`
	RegIp     string    `json:"reg_ip,omitempty"`
}
