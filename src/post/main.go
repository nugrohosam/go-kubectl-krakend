package main

import (
	"os"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(healthcheck.Default())

	r.POST("/level1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	r.POST("/level1/level2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	r.POST("/level1/level2/level3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
