package main

import (
	"os"
	"strconv"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pagesize := os.Getpagesize()

	r := gin.Default()
	r.Use(healthcheck.Default())
	r.GET("/get-ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
