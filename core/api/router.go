package api

import "github.com/labstack/echo"

// Router init router
func Router(e *echo.Echo) {
	e.GET("/", Ping)
}
