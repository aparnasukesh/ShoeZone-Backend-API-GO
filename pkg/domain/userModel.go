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

type Address struct {
	gorm.Model
	Street      string `json:"street" validate:"required,max=255"`
	City        string `json:"city" validate:"required,max=255"`
	State       string `json:"state" validate:"required ,max=255"`
	PINCode     uint   `json:"pin_code" validate:"required,min=6,max=6"`
	Country     string `json:"country" validate:"max=100"`
	UserID      int    `json:"user_id"`
	UserAddress User   `gorm:"foreignkey:UserID"`
}

// Create a new struct for updating user profile without Isverified, Isadmin, and Otp fields
type UserProfileUpdate struct {
	Username    string `json:"username" validate:"required,min=8,max=24"`
	Password    string `json:"password" validate:"required,min=8,max=16"`
	Phone       string `json:"phone" validate:"required,len=10"`
	Email       string `json:"email" validate:"email,required"`
	Dateofbirth string `json:"dateofbirth"`
	Gender      string `json:"gender"`
}

type ResetPassword struct {
	OTP         string `json:"otp"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=16"`
}
