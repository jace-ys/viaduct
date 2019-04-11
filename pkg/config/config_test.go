package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "3000")
	os.Setenv("CONFIG_PATH", "config/config.sample.yml")

	conf, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "3000", conf.Port)
	assert.Equal(t, "config/config.sample.yml", conf.ConfigPath)
}

func TestDefineServices(t *testing.T) {
	serviceRegistry, err := RegisterServices("../../config/config.sample.yml")
	if err != nil {
		t.Error(err)
	}

	testService := serviceRegistry.Services["test"]

	assert.Equal(t, "Testing", testService.Name)
	assert.Equal(t, "test", testService.Prefix)
	assert.Equal(t, "http://testing.com", testService.UpstreamUrl.String())
	assert.Equal(t, "GET", testService.Methods[0])
	assert.Equal(t, true, testService.AllowCrossOrigin)
	assert.Equal(t, true, testService.Middlewares["logging"])
}
