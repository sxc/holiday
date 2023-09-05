package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type TimeoffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"required,future" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

var ValidatorFuture validator.Func = func(f1 validator.FieldLevel) bool {
	date, ok := f1.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("future", ValidatorFuture)
	}

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	router.POST("/employee", func(c *gin.Context) {
		var timeoffRequest TimeoffRequest
		if err := c.ShouldBind(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	apiGroup := router.Group("/api")
	apiGroup.POST("/timeoff", func(c *gin.Context) {
		var timeoffRequest TimeoffRequest
		if err := c.ShouldBind(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	router.StaticFile("/", "./public/index.html")

	router.GET("/tale_of_two_cities", func(c *gin.Context) {
		f, err := os.Open("./tale_of_two_cities.txt")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()
		data, err := io.ReadAll(f)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Data(http.StatusOK, "text/plain", data)

	})

	router.GET("/grate_expectations", func(c *gin.Context) {
		f, err := os.Open("./grate_expectations.txt")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.DataFromReader(http.StatusOK,
			fi.Size(),
			"text/plain",
			f,
			map[string]string{
				"Content-Disposition": `attachment; filename="tale_of_two_cities.txt"`,
			},
		)
	})

	log.Fatal(router.Run(":3000"))
}
