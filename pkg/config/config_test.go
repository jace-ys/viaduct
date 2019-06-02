package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "3000")
	os.Setenv("CONFIG_FILE", "config/config.sample.yaml")

	conf, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "3000", conf.Port)
	assert.Equal(t, "config/config.sample.yaml", conf.ConfigFile)
}
