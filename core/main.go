package main

import (
	"quote-board/core/api"
	"quote-board/core/models"
	"quote-board/core/settings"

	"github.com/labstack/echo"
)

func main() {
	setting := settings.ServerSetting{}
	setting.ParseEnv()

	models.InitDB(setting.Postgres)

	e := echo.New()
	e.Debug = setting.Debug

	api.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
