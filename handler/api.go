package handler

import (
	"github.com/gin-gonic/gin"

	"kappa-web/pkg/res"
)

func Ping(c *gin.Context) {
	res.Success(c, gin.H{
		"message": "pong",
	})
}
