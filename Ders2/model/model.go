package model

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"unique;column:user_name"`
	Password string	`gorm:"size:255;column:password"`
}

func (u User) TableName() string {
	return "users"
}
