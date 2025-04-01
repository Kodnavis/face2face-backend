package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Login     string `gorm:"unique" json:"login"`
	Password  string `json:"password"`
}
