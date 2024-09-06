package models

import (
	"gorm.io/gorm"
)

type SignupUser struct {
	gorm.Model
	Phone         string `gorm:"unique" json:"phone" bson:"phone"`
	Name          string `json:"name" bson:"name"`
	Password      string `json:"password" bson:"password"`
	ConfirmPassword string `json:"confirmPassword" bson:"confirmPassword"`
	Address       string `json:"address" bson:"address"`
	DriverLicense string `json:"driverLincense" bson:"driverLicense"`
	FaceCapture   string `json:"faceCapture" bson:"faceCapture"`
}

type LoginUser struct {
	Phone       string `json:"phone" bson:"phone"`
	Password    string `json:"password" bson:"password"`
	FaceCapture string `json:"faceCapture" bson:"faceCapture"`
}
