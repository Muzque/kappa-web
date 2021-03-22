package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"kappa-web/handler"
	"kappa-web/config"
)

func main() {
	log.Println("Initializing ApiServer...")
	load()

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	router.GET("/ping", handler.Ping)

	api := router.Group("/api")
	{
		api.GET("/ping", handler.Ping)
	}
	router.Run(":" + config.Val.Port)
}


func load() {
	config.Init()
}