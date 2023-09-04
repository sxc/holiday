package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// Static Routes
	router.Static("/statics", "./statics")

	// Static Routes to string
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	log.Fatal(router.Run(":3000")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
