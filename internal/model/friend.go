package model

import "gorm.io/gorm"

type Friend struct {
	*gorm.Model
	FromUser  uint
	ToUser    uint `json:"id,omitempty"`
	Confirmed bool
}
