package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err: %v\n", err)
	}
}
