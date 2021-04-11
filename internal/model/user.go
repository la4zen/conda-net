package model

import "time"

type User struct {
	ID        int        `json:"id,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Login     string     `json:"login,omitempty"`
	Password  string     `json:"password,omitempty"`
	RegDate   *time.Time `json:"reg_date,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	RegIp     string     `json:"reg_ip,omitempty"`
}
