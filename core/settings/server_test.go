package settings

import "testing"

func TestServerSetting_ParseEnv(t *testing.T) {
	tests := []struct {
		name string
		s    *ServerSetting
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.ParseEnv()
		})
	}
}
