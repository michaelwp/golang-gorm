package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50); not null"`
	LastName string `gorm:"type:varchar(50); not null"`
	Credential Credential
}

type Credential struct {
	gorm.Model
	UserID uint
	Email string `gorm:"unique; not null"`
	Password string `gorm:"not null"`
}

type UserInput struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Result struct {
	 Status int `json:"status"`
	 Message string `json:"message"`
	 Data []User
}