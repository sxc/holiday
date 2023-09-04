package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()

	// router.StaticFile("/", "./public/index.html")

	// router.Static("/public", "./public")

	// router.StaticFS("/fs", http.FileSystem(http.FS(f)))

	// http.Request object
	// router.GET("/*rest", func(c *gin.Context) {
	// 	url := c.Request.URL.String()
	// 	headers := c.Request.Header
	// 	cookies := c.Request.Cookies()

	// 	c.IndentedJSON(http.StatusOK, gin.H{
	// 		"url":     url,
	// 		"headers": headers,
	// 		"cookies": cookies,
	// 	})
	// })

	// http://localhost:3000/query/?username=john&year=2010&month=1&month=2

	router.GET("/query/*rest", func(c *gin.Context) {
		username := c.Query("username")
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := c.QueryArray("month")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		date := c.PostForm("date")
		amount := c.PostForm("amount")
		username := c.DefaultPostForm("username", "john")

		c.IndentedJSON(http.StatusOK, gin.H{
			"date":     date,
			"amount":   amount,
			"username": username,
		})
	})

	log.Fatal(router.Run(":3000"))
}
