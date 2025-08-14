package main

import (
	"CloudCast/Config"
	"CloudCast/Routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi DB pakai pgx.Conn
	connectDB, err := Config.CloudCastDBConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Connected to database:", connectDB.Config().ConnString())

	// Setup router
	r := gin.Default()
	r.LoadHTMLGlob("Views/*")

	// Register semua route dengan DB
	Routes.RegisterRoutes(r, connectDB)

	// Jalankan server
	r.Run(":8080")
}
