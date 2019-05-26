package api

import "github.com/labstack/echo"

func Router(e *echo.Echo) {
	e.GET("/", Ping)
}
