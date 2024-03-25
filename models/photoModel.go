package models

import (
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Photo struct {
	ID string
	Title string
	Caption string
	PhotroUrl string
	UserId uint
	User *User
	CreatedAt string
	UpdatedAt string
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}