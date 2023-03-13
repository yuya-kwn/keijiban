package middleware

import (
	"keijiban/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contian an access token"})
			c.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}