package main

import (
	"github.com/codingno/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	routes.MainRoute = gin.Default()

	routes.MainRoute.Run(":3000")
}
