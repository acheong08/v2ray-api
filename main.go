package main

import (
	"os"

	"github.com/acheong08/v2ray-api/trojan"
	"github.com/gin-gonic/gin"
)

func admin_auth(c *gin.Context) {
	// Get Authorization header
	auth_header := c.GetHeader("Authorization")
	// Check if the header matches env variable
	if auth_header == os.Getenv("ADMIN_AUTH") {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}

func main() {

	tr := trojan.Trojan{}

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	server.POST("/admin/start", admin_auth, func(c *gin.Context) {
		tr.Start()
		c.JSON(200, gin.H{"message": "started"})
	})

	server.POST("/admin/stop", admin_auth, func(c *gin.Context) {
		tr.Stop()
		c.JSON(200, gin.H{"message": "stopped"})
	})

	// Run
	server.Run(":8080")

}
