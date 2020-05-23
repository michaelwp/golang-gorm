package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50); not null" json:"first_name"`
	LastName string `gorm:"type:varchar(50); not null" json:"last_name"`
	Credential Credential
}

type Credential struct {
	gorm.Model
	Email string `gorm:"unique; not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
