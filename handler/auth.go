package handler

import (
	"github.com/gin-gonic/gin"
	"kappa-web/pkg/jwt"
	"kappa-web/pkg/res"
)

type LoginForm struct {
	Account string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context)  {
	form := LoginForm{}
	err := c.BindJSON(&form)
	if err != nil {
		res.BadRequest(c, 403, gin.H{
			"message": "Fail to parse user account.",
		})
		return
	}

	token, err := jwt.GenerateToken(form.Account, form.Password)
	if err != nil {
		res.BadRequest(c, 403, gin.H{
			"message": "Fail to generate token",
		})
		return
	}

	res.Success(c, gin.H{
		"token": token,
	})
}
