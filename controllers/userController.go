package controllers

import (
	"assignment4_test/models"
	"assignment4_test/config"
	_"assignment4_test/database"
	"assignment4_test/helpers"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	_"github.com/asaskevich/govalidator"

	"time"
	_"fmt"
	"net/http"
)

var userDatas = []models.User{}

func CreateUser(c *gin.Context) {
	db := config.GetDB()

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == "application/JSON" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID =  uuid.New().String()
	User.CreatedAt = time.Now().String()
	User.UpdatedAt = time.Now().String()


	// _, errCreate := govalidator.ValidateStruct(&User)

	// if errCreate != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "bad request",
	// 		"message": errCreate.Error(),
	// 	})
	// 	return
	// }

	if User.Email == "" || User.Password == "" || User.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": "email or password or username can't be empty",
		})
		return
	}

	if User.Age < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": "minimal age is 8",
		})
		return
	}

	pw := []rune(User.Password)
	if len(pw) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": "password too short",
		})
		return
	}

	if errem := helpers.EmailValidator(User.Email); errem !=  nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": "invalid email",
		})
		return
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age": User.Age,
		"email": User.Email,
		"id": User.ID,
		"username": User.UserName,

	})
	// var newUser models.User

	// if err := c.ShouldBindJSON(&newUser); err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }



	// errdb := database.CreateUser(newUser.Email, newUser.Password, newUser.Age, newUser.UserName)

	// if errdb != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": errdb,
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusCreated, gin.H{
	// 	"age": newUser.Age,
	// 	"email": newUser.Email,
	// 	"id": newUser.ID,
	// 	"username": newUser.UserName,
	// })

}

// func LoginUser(c *gin.Context) {
// 	var login models.User

// 	if err := c.ShouldBindJSON(&login); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	test := database.GetByEmail(login.Email)

// 	c.JSON(http.StatusCreated, gin.H{
// 		"login": test,
// 	})

// 	// for _, user := range userDatas {
// 	// 	if login.Email == user.Email {
// 	// 		if login.Password == user.Password {
// 	// 			c.JSON(http.StatusCreated, gin.H{
// 	// 				"success" : true,
// 	// 			})
// 	// 		} else {
// 	// 			c.JSON(http.StatusCreated, gin.H{
// 	// 				"success": "password invalid",
// 	// 			})
// 	// 		}
// 	// 		break
// 	// 	} else {
// 	// 		c.JSON(http.StatusCreated, gin.H{
// 	// 			"success": "user not found",
// 	// 		})
// 	// 	}
// 	// }
// }

func UserLogin(c *gin.Context) {
	db := config.GetDB()

	contentType := helpers.GetContentType(c)

	User := models.User{}

	if contentType == "application/JSON" {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"JWT": token,
	})
}
