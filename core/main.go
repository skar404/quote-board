package main

import (
	"github.com/labstack/echo"
	"sharelink/core/api"
	"sharelink/core/models"
	"sharelink/core/settings"
)

func main() {
	setting := settings.ServerSetting{
		Postgres: "postgres://postgres:postgres@localhost:5400/postgres?sslmode=disable",
	}

	models.InitDB(setting.Postgres)

	e := echo.New()
	e.GET("/", api.Ping)
	e.Logger.Fatal(e.Start(":8080"))
}
