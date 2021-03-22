package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"kappa-web/handler"
)

func NotFound(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Homepage",
	})
}

func main() {
	log.Println("Initializing ApiServer...")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	router.GET("/ping", handler.Ping)

	api := router.Group("/api")
	{
		api.GET("/ping", handler.Ping)
	}
	router.Run(":7777")
}
