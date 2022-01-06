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
	r.GET("/level1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	r.GET("/level1/level2/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	r.GET("/level1/level2/level3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
