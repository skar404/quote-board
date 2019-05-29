package models

import "testing"

func TestPingDB(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PingDB(); (err != nil) != tt.wantErr {
				t.Errorf("PingDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
