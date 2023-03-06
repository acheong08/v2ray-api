package main

import (
	"github.com/acheong08/v2ray-api/trojan"
	"github.com/gin-gonic/gin"
)

var trojan_server *trojan.Trojan

func init() {
	trojan_server = trojan.New()
}

func main() {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	server.POST("/restart", restart)

	server.POST("/stop", stop)

	// Run
	server.Run(":8080")

}

// Handlers
func restart(c *gin.Context) {
	// Get JSON body as string
	config, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	// Create trojan server
	if trojan_server.Status() == "exists" {
		err = trojan_server.RestartWithNewConfig(string(config))
	} else {
		err = trojan_server.CreateAndRun(string(config))

	}
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}
}

func stop(c *gin.Context) {
	err := trojan_server.Stop()
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"message": "Stopped"})
}
