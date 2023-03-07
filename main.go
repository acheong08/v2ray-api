package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/acheong08/v2ray-api/trojan"
	"github.com/gin-gonic/gin"
)

func authMiddleware(expectedToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		actualToken := c.GetHeader("Authorization")
		if actualToken != expectedToken {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func main() {
	tr := trojan.Trojan{}
	// Start by default
	if err := tr.Start(); err != nil {
		panic(err)
	}
	router := gin.Default()

	// Ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Admin routes
	adminGroup := router.Group("/admin")
	adminGroup.Use(authMiddleware(os.Getenv("ADMIN_AUTH")))
	{
		// Start route
		adminGroup.POST("/start", func(c *gin.Context) {
			if tr.Status() == "running" {
				c.JSON(http.StatusOK, gin.H{"message": "already running"})
				return
			}
			if err := tr.Start(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "started"})
		})

		// Stop route
		adminGroup.POST("/stop", func(c *gin.Context) {
			if tr.Status() == "stopped" {
				c.JSON(http.StatusOK, gin.H{"message": "already stopped"})
				return
			}
			if err := tr.Stop(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "stopped"})
		})

		// Restart route
		adminGroup.POST("/restart", func(c *gin.Context) {
			if err := tr.Restart(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "restarted"})
		})

		// Status route
		adminGroup.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": tr.Status()})
		})

		// Configure route
		adminGroup.POST("/configure", func(c *gin.Context) {
			var config interface{}
			if err := c.ShouldBindJSON(&config); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
				return
			}
			jsonConfig, err := json.Marshal(config)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			if err := tr.Configure(string(jsonConfig)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "configured"})
		})

		// Config route
		adminGroup.GET("/config", func(c *gin.Context) {
			config, err := tr.GetConfig()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			var jsonConfig interface{}
			if err := json.Unmarshal([]byte(config), &jsonConfig); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, jsonConfig)
		})
	}
	// Run
	router.Run()
}
