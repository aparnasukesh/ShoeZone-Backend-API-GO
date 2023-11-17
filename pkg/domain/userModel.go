package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" validate:"required,min=8,max=24"`
	Password    string `json:"password" validate:"required,min=8,max=16"`
	Phone       string `json:"phone" validate:"required,len=10"`
	Email       string `json:"email" validate:"email,required"`
	Otp         string `json:"otp"`
	Isverified  bool   `json:"isverified" gorm:"default:true"`
	Isadmin     bool   `json:"isadmin" gorm:"default:false"`
	Dateofbirth string `json:"dateofbirth"`
	Gender      string `json:"gender"`
}
