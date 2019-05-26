package api

import (
	"github.com/labstack/echo"
	"net/http"
	"sharelink/core/models"
)

type pingView struct {
	Status string `json:"status"`
	Text   string `json:"text,omitempty"`
}

// Ping ping server
func Ping(c echo.Context) error {
	err := models.PingDB()
	if err != nil {
		return c.JSON(http.StatusOK, pingView{Status: "ERROR", Text: "not connect to db"})
	}

	return c.JSON(http.StatusOK, pingView{Status: "SUCCESS"})
}
