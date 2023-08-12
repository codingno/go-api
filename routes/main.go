package routes

import "github.com/gin-gonic/gin"

var MainRoute *gin.Engine

func NewRoute(method string, path string, handler gin.HandlerFunc) {
	for {
		if MainRoute != nil {
			break
		}
	}
	MainRoute.Handle(method, path, handler)
}
