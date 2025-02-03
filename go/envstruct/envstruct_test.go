package envstruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	App AppConfig
	Web WebConfig
}

type AppConfig struct {
	Enviroment string `default:"one"`
	Version    int
	Number     string
}

type WebConfig struct {
	Http struct {
		Server struct {
			Port string
			Name string
			Some int `default:"1"`
		}
	}
	Constrains struct {
		Some int    `env:"int_special_key"`
		Msg  string `env:"special_key"`
	}
}

func TestParse(t *testing.T) {
	var err error
	a := Config{}
	err = Parse(&a)
	assert.NoError(t, err)
	assert.Equal(t, "8080", a.Web.Http.Server.Port)
	assert.Equal(t, 10, a.App.Version)
}

func TestParse_With_DefaultTag(t *testing.T) {
	var err error
	a := Config{}
	err = Parse(&a)
	assert.NoError(t, err)
	assert.Equal(t, "one", a.App.Enviroment)
	assert.Equal(t, 1, a.Web.Http.Server.Some)
}

func TestParse_With_EnvTag(t *testing.T) {
	var err error
	a := Config{}
	err = Parse(&a)
	assert.NoError(t, err)
	assert.Equal(t, "servus", a.Web.Constrains.Msg)
	assert.Equal(t, 20, a.Web.Constrains.Some)
}

func TestPars_With_Preterm(t *testing.T) {
	var err error
	a := AppConfig{}
	Preterm = "MASLIAS"
	err = Parse(&a)
	assert.NoError(t, err)
	assert.Equal(t, "30", a.Number)
	assert.Equal(t, 50, a.Version)
}

func TestParse_Invalid_Struct(t *testing.T) {
	var err error
	err = Parse(nil)
	assert.Error(t, err)

	c := Config{}
	err = Parse(c)
	assert.Error(t, err)

	err = Parse(66)
	assert.Error(t, err)

	var ptr bool
	err = Parse(ptr)
	assert.Error(t, err)
}
