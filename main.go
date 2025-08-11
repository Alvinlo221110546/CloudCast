package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Cuaca siap digunakan!",
		})
	})

	r.Run(":8080") // Jalankan server di port 8080
}
