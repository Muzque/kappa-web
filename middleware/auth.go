package middleware

import (
	"github.com/gin-gonic/gin"
	"kappa-web/pkg/jwt"
	"kappa-web/pkg/res"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(jwt.Key)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error_code": res.ErrUnauthorizedCode,
			})
			c.Abort()
			return
		}

		account, err := jwt.VerifyToken(token)
		if err != nil || account == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error_code": res.ErrUnauthorizedCode,
			})
			c.Abort()
			return
		}
		c.Set("account", account)
		c.Next()
	}
}
