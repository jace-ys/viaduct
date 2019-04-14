package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "3000")
	os.Setenv("CONFIG_FILE", "config/config.sample.yml")

	conf, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "3000", conf.Port)
	assert.Equal(t, "config/config.sample.yml", conf.ConfigFile)
}

func TestDefineApis(t *testing.T) {
	apiRegistry, err := RegisterApis("../../config/config.sample.yml")
	if err != nil {
		t.Error(err)
	}

	testApi := apiRegistry.Apis["test"]

	assert.Equal(t, "Testing", testApi.Name)
	assert.Equal(t, "test", testApi.Prefix)
	assert.Equal(t, "http://localhost:8080/invalid/url", testApi.UpstreamUrl.String())
	assert.Equal(t, "GET", testApi.Methods[0])
	assert.Equal(t, true, testApi.AllowCrossOrigin)
	assert.Equal(t, true, testApi.Middlewares["logging"])
}
