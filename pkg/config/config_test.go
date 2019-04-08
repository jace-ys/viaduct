package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "80", conf.Port)
	assert.Equal(t, "/config/config.yml", conf.FilePath)
	assert.Equal(t, "viaduct", conf.LogPrefix)
}

func TestDefineServices(t *testing.T) {
	serviceDefinitions, err := RegisterServices("../../config/config.sample.yml")
	if err != nil {
		t.Error(err)
	}

	testService := serviceDefinitions.Services["test"]

	assert.Equal(t, "Testing", testService.Name)
	assert.Equal(t, "test", testService.Prefix)
	assert.Equal(t, "http://testing.com", testService.UpstreamUrl.String())
	assert.Equal(t, "GET", testService.Methods[0])
	assert.Equal(t, true, testService.AllowCrossOrigin)
}
