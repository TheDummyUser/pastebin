package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/thedummyuser/pastebin/db"
	"github.com/thedummyuser/pastebin/helpers"
	"github.com/thedummyuser/pastebin/routes"
)

func main() {
	database, err := db.InitDb()

	if err != nil {
		fmt.Print("somethign went wront while migrating db", err)
	}

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	routes.RegisterRoutes(e, database)

	e.Logger.Fatal(e.Start(":1234"))
}
