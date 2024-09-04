package models

import (
	"gorm.io/gorm"
)

type SignupUser struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email" bson:"email"`
	Phone    string `gorm:"unique" json:"phone" bson:"phone"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type LoginUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
