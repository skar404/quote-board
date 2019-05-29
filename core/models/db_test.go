package models

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestInitDB(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDB(tt.args.dataSourceName)
		})
	}
}
