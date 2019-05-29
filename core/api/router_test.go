package api

import (
	"testing"

	"github.com/labstack/echo"
)

func TestRouter(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Router(tt.args.e)
		})
	}
}
