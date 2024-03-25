package routers

import (
	"assignment4_test/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.CreateUser)
	r.POST("/users/login", controllers.UserLogin)


	return r
}