package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestMain(m *testing.M) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")

	if assert.NoError(t, h.getUser(c)) {
		text := rec.Body.String()
		fmt.Print(text)
	}
}
