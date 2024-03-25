package models

import (
	"gorm.io/gorm"
	// "github.com/asaskevich/govalidator"
	"assignment4_test/helpers"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"not null;uniqueIndex" valid:"requred-Your Username is required"`
	Email     string `gorm:"not null;uniqueIndex" valid:"required-Your Email is required,email-Invalid email format"`
	Password  string `gorm:"not null" valid:"required-Your password is required,minstringlength(6)-Password has to have a minimum length of 6 characters"`
	Age       int `gorm:"not null" valid:"required-Your Age is requred"`
	CreatedAt string
	UpdatedAt string
	Photos []Photo
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	// _, errCreate := govalidator.ValidateStruct(u)

	// if errCreate != nil {
	// 	err = errCreate
	// 	return err
	// }

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return err
}