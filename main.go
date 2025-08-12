package main

import (
	// "context"
	// "fmt"
	// "os"

	"CloudCast/Config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	connectDB, err := Config.CloudCastDBConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Connected to database:", connectDB.Config().ConnString())

	r := gin.Default()
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
