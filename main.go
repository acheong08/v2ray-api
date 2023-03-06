package main

import (
	"encoding/json"
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
		if tr.Status() == "running" {
			c.JSON(200, gin.H{"message": "already running"})
			return
		}
		err := tr.Start()
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "started"})
	})

	server.POST("/admin/stop", admin_auth, func(c *gin.Context) {
		if tr.Status() == "stopped" {
			c.JSON(200, gin.H{"message": "already stopped"})
			return
		}
		err := tr.Stop()
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "stopped"})
	})

	server.POST("/admin/restart", admin_auth, func(c *gin.Context) {
		err := tr.Restart()
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "restarted"})
	})

	server.GET("/admin/status", admin_auth, func(c *gin.Context) {
		c.JSON(200, gin.H{"status": tr.Status()})
	})

	server.POST("/admin/configure", admin_auth, func(c *gin.Context) {
		var config interface{}
		c.BindJSON(&config)
		// Convert config to JSON string
		json_config, err := json.Marshal(config)
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		err = tr.Configure(string(json_config))
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "configured"})
	})

	// Run
	server.Run(":8080")

}
