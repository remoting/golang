package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	fmt.Println("Web Server")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
	fmt.Println("listen and serve on 0.0.0.0:8080")
}
