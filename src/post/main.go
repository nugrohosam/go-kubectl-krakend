package main

import (
	"fmt"
	"os"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	nats "github.com/nats-io/nats.go"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(healthcheck.Default())

	nc, _ := nats.Connect("nats://" + os.Getenv("NATS_HOST"))

	r.POST("/level1", func(c *gin.Context) {
		nc.Publish("from-post", []byte("Hello World level1"))
		fmt.Println("published")
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	r.POST("/level1/level2", func(c *gin.Context) {
		nc.Publish("from-post", []byte("Hello World level12"))
		fmt.Println("published")
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	r.POST("/level1/level2/level3", func(c *gin.Context) {
		nc.Publish("from-post", []byte("Hello World level123"))
		fmt.Println("published")
		c.JSON(200, gin.H{
			"message": "pong something from name : " + name})
	})

	nc.Subscribe("from-get", func(m *nats.Msg) {
		fmt.Printf("Received a message in post: %s\n", string(m.Data))
	})

	r.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
