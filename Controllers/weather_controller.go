package controllers

import (
	"CloudCast/Models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func HomeHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := Models.GetWeather(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		tmpl, _ := template.ParseFiles("Views/Home.html")
		tmpl.Execute(w, data)
	}
}

// func WeatherHandler(db *pgx.Conn) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data, err := Models.GetWeather(db)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		tmpl, _ := template.ParseFiles("Views/Weather.html")
// 		tmpl.Execute(w, data)
// 	}
// }

// func AboutHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		tmpl, _ := template.ParseFiles("Views/About.html")
// 		tmpl.Execute(w, nil)
// 	}
// }

func WrapHTTPHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c.Writer, c.Request)
	}
}
