package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello, World!"})
	})

	r.GET("/haha/:id", func(c *gin.Context) {
		var a = c.Param("id")
		c.JSON(http.StatusOK, gin.H{"data": a})
	})

	r.Run()
}
