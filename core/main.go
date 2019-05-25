package main

import (
	"github.com/labstack/echo"

	"sharelink/core/api"
)

func main() {

	e := echo.New()
	e.GET("/", api.Ping)
	e.Logger.Fatal(e.Start(":8080"))
}
