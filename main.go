package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// Static Routes
	router.StaticFile("/", "./public/index.html")

	router.Static("/public", "./public")

	router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	// Static Routes to string
	router.GET("/hello/world", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	log.Fatal(router.Run(":3000")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
