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
	r.POST("/post-ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from here name : " + name,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
