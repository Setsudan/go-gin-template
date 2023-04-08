package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"timestamp": time.Now().Unix(),
		})
	})
}
