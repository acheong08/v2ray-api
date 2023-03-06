package main

import (
	"github.com/acheong08/v2ray-api/trojan"
	"github.com/gin-gonic/gin"
)

func main() {

	tr := trojan.Trojan{}

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	server.POST("/start", func(c *gin.Context) {
		tr.Start()
		c.JSON(200, gin.H{"message": "started"})
	})

	server.POST("/stop", func(c *gin.Context) {
		tr.Stop()
		c.JSON(200, gin.H{"message": "stopped"})
	})

	// Run
	server.Run(":8080")

}
