package api

import (
	"net/http"

	"quote-board/core/models"

	"github.com/labstack/echo"
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
