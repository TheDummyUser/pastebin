package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thedummyuser/pastebin/routes/handlers"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/:uuid", func(c echo.Context) error {
		return handlers.GetSinglePost(c, db)
	})

	e.POST("/paste", func(c echo.Context) error {
		return handlers.Paste(c, db)
	})
}
