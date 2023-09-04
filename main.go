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

	router.StaticFile("/", "./public/index.html")

	router.Static("/public", "./public")

	router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	// Static Routes to string
	router.GET("/hello/world", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	// Routing with HTTP verbs
	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		c.String(http.StatusOK, "New request POSTed successfully")
	})

	// var username string
	// Parameterized Routes
	router.GET("/employees/:username/*rest", func(c *gin.Context) {
		username := c.Param("username")
		rest := c.Param("rest")
		wholeroute := c.FullPath()

		c.String(http.StatusOK, "Username: "+username+", Rest: "+rest+", Full route: "+wholeroute)

	})

	log.Fatal(router.Run(":3000")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
