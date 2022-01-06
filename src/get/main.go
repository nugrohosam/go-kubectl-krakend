package main

import (
	"fmt"
	"os"
	"strconv"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	nats "github.com/nats-io/nats.go"
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

	nc, _ := nats.Connect("nats://nats-service:4222")

	pagesize := os.Getpagesize()

	r := gin.Default()
	r.Use(healthcheck.Default())
	r.GET("/level1", func(c *gin.Context) {
		nc.Publish("from-get", []byte("Hello World level1"))
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	r.GET("/level1/level2/", func(c *gin.Context) {
		nc.Publish("from-get", []byte("Hello World level12"))
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	r.GET("/level1/level2/level3", func(c *gin.Context) {
		nc.Publish("from-get", []byte("Hello World level123"))
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name + " in " + pwd + " pagesize " + strconv.Itoa(pagesize) + " woow ",
		})
	})

	// Simple Publisher
	nc.Subscribe("from-post", func(m *nats.Msg) {
		fmt.Printf("Received a message in get: %s\n", string(m.Data))
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
