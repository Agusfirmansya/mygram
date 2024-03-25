package helpers

// import (
// 	"assignment4_test/models"
// 	"fmt"
// 	"time"

// 	"encoding/json"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// )

// var secretKey = []byte("secret-key")

// func CreateToken(email string) (string, error){
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
//         jwt.MapClaims{ 
//         "username": email, 
//         "exp": time.Now().Add(time.Hour * 24).Unix(), 
//         })

//     tokenString, err := token.SignedString(secretKey)
//     if err != nil {
//     return "", err
//     }

//  	return tokenString, nil
// }

// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 	   return secretKey, nil
// 	})
   
// 	if err != nil {
// 	   return err
// 	}
   
// 	if !token.Valid {
// 	   return fmt.Errorf("invalid token")
// 	}
   
// 	return nil
// }

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"errors"
	"strings"
)

var secretKey = "secret"

func GenerateToken(id string, email string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign in to proceed")

	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}