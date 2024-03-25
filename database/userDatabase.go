package database

import (
	"fmt"
	"time"
	"github.com/google/uuid"

	"assignment4_test/config"
	"assignment4_test/models"
	"assignment4_test/helpers"
)


func CreateUser(email string, password string, age int, username string) error{
	db := config.GetDB()

	if age < 8 {
		return fmt.Errorf("minimal age is 8")
	}

	if errem := helpers.EmailValidator(email); errem !=  nil {
		return errem
	}
	
	User := models.User{
		Email: email,
		Password: password,
		Age: age,
		UserName: username,
		ID:  uuid.New().String(),
		CreatedAt: time.Now().String(), //fmt.Sprintf("%s", time.Now()),
		UpdatedAt: time.Now().String(), //fmt.Sprintf("%s", time.Now()),
	}

	err := db.Create(&User).Error

	if err != nil {
		return err
	}

	return nil

	//fmt.Println(User)
}

func GetByEmail(email string) models.User{
	// var productRead Product
    // db.First(&productRead, 1)
	db := config.GetDB()
	var getUser models.User

	db.First(&getUser, "email = ?", email)

	return getUser
}