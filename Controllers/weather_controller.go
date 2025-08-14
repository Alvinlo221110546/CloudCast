package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"CloudCast/Models"
)

func HomeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := Models.GetWeather(db)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	}
}
