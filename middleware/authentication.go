package middleware

import {
	"github.com/gin-gonic/gin"
	"assignment4_test/helpers"
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AborthWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}