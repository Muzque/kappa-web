package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"kappa-web/config"
	"kappa-web/handler"
	"kappa-web/middleware"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	router.GET("/ping", handler.Ping)

	api := router.Group("/api")
	{
		api.GET("/ping", middleware.Auth(), handler.Ping)
	}
	return router
}


func main() {
	log.Println("Initializing ApiServer...")
	load()
	router := setupRouter()
	err := router.Run(":" + config.Val.Port)
	if err != nil {
		return 
	}
}


func load() {
	config.Init()
}