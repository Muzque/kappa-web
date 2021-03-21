package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing ApiServer...")
	router := gin.Default()
	router.LoadHTMLGlob("kappa/dist/*.html")
	router.Static("/js", "kappa/dist/js")
	router.Static("/css", "kappa/dist/css")
	router.Static("/img", "kappa/dist/img")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Kappa",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Kappa",
		})
	})

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":7777")
}
