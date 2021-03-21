package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Homepage",
	})
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Homepage",
	})
}

func main() {
	log.Println("Initializing ApiServer...")
	router := gin.Default()
	router.LoadHTMLGlob("kappa/dist/*.html")
	router.Static("/js", "kappa/dist/js")
	router.Static("/css", "kappa/dist/css")
	router.Static("/img", "kappa/dist/img")

	router.GET("/", Home)
	router.GET("/ping", Ping)
	router.NoRoute(NotFound)
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", Ping)
	}
	router.Run(":7777")
}
