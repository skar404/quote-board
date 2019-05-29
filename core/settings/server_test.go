package settings

import (
	"log"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
)

func TestServerSetting_ParseEnv(t *testing.T) {
	opts := ServerSetting{}

	var parser = flags.NewParser(&opts, flags.Default)
	_, err := parser.ParseArgs([]string{})

	if err != nil {
		log.Panic("not parse env", err)
	}

	assert.Equal(t, opts.Postgres, "postgres://postgres:postgres@localhost:5400/postgres?sslmode=disable")
	assert.False(t, opts.Debug)
}
