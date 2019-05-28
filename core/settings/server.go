package settings

import (
	"github.com/jessevdk/go-flags"
	"github.com/labstack/gommon/log"
)

// ServerSetting base app settings or env
type ServerSetting struct {
	Postgres string `long:"postgres" env:"POSTGRES" default:"postgres://postgres:postgres@localhost:5400/postgres?sslmode=disable"`
	Debug    bool   `long:"debug" env:"DEBUG"`
}

// ParseEnv parse env and set default value
func (s *ServerSetting) ParseEnv() {
	var parser = flags.NewParser(s, flags.Default)
	_, err := parser.Parse()

	if err != nil {
		log.Panic("not parse env", err)
	}
}
