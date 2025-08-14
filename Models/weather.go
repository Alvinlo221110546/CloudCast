package Models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type WeatherToday struct {
	ID          int
	City        string
	Temperature int
	Condition   string
	Humidity    int
	WindSpeed   int
	Date        string
}

func InsertWeather(db *pgx.Conn, w WeatherToday) error {
	_, err := db.Exec(
		context.Background(),
		`INSERT INTO weather_today (city, temperature, condition, humidity, wind_speed)
		 VALUES ($1, $2, $3, $4, $5)`,
		w.City, w.Temperature, w.Condition, w.Humidity, w.WindSpeed,
	)
	return err
}

func GetWeather(db *pgx.Conn) ([]WeatherToday, error) {
	rows, err := db.Query(
		context.Background(),
		`SELECT id, city, temperature, condition, humidity, wind_speed, date FROM weather_today`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weathers []WeatherToday
	for rows.Next() {
		var w WeatherToday
		if err := rows.Scan(&w.ID, &w.City, &w.Temperature, &w.Condition, &w.Humidity, &w.WindSpeed, &w.Date); err != nil {
			return nil, err
		}
		weathers = append(weathers, w)
	}
	return weathers, nil
}
