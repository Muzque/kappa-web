package middleware

import (
	"github.com/gin-gonic/gin"
	"kappa-web/pkg/jwt"
	"kappa-web/pkg/res"
)

func unauthorized(c *gin.Context)  {
	res.BadRequest(c, 403, gin.H{
		"message": "Unauthenticated credentials.",
	})
	c.Abort()
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(jwt.Key)
		if err != nil {
			unauthorized(c)
			return
		}

		account, err := jwt.VerifyToken(token)
		if err != nil || account == "" {
			unauthorized(c)
			return
		}
		c.Set("account", account)
		c.Next()
	}
}
