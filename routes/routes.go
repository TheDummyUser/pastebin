package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thedummyuser/pastebin/routes/handlers"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {

	e.GET("/getpasts", func(c echo.Context) error {
		return handlers.GetAllPastes(c, db)
	})

	e.POST("/paste", func(c echo.Context) error {
		return handlers.Paste(c, db)
	})
}
