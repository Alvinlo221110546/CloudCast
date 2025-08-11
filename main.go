package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set folder views
	r.LoadHTMLGlob("Views/*")

	// Route Home
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Home.html", nil)
	})

	// Route Weather
	r.GET("/weather", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Weather.html", nil)
	})

	// Route About
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "About.html", nil)
	})

	// Jalankan server di port 8080
	r.Run(":8080")
}
