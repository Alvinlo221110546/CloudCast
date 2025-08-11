package main

import (
	"context"
	"fmt"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	databaseUrl := "postgres://postgres:alpintod@localhost:5432/weathercast"
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	fmt.Println("Successfully connected to PostgreSQL!")

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
