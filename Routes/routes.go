package Routes

import (
	controllers "CloudCast/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutes(r *gin.Engine, db *pgx.Conn) {
	r.GET("/", controllers.WrapHTTPHandler(controllers.HomeHandler(db)))
	// r.GET("/weather", controllers.WrapHTTPHandler(controllers.WeatherHandler(db)))
	// r.GET("/about", controllers.WrapHTTPHandler(controllers.AboutHandler()))
}
