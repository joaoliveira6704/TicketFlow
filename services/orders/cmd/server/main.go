package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var orders = []string{"order1", "order2", "order3"}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func listOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": orders})
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/orders", listOrders)
		api.GET("/health", health)
	}

	router.Run(":8080")
}
