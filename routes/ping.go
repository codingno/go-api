package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	go NewRoute("GET", "/ping", PingHandler)
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
