package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		h, _ := os.Hostname()
		c.JSON(200, gin.H{
			"hostname": h,
		})
	})
	r.Run(":8080")
}
