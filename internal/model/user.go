package model

type User struct {
	ID        int    `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func (u *User) Exists() bool {
	return false
}
