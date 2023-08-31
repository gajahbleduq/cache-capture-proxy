package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/tomimulhartono/cache-capture-proxy/handler"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.ProxyHandler)

	router.Run(":8080")
}
