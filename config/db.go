package config

import (
	"assignment4_test/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "postgres"
	password = "123"
	dbPort = "5432"
	dbname = "postgres"
	db *gorm.DB
	err error
)
func ConnectDatabase() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort) 
	//"host=localhost user=postgres dbname=postgres sslmode=disable"


	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	
	db.Debug().AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}
