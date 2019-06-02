package config

import (
	"github.com/caarlos0/env"
)

var (
	DefaultPort       = "80"
	DefaultConfigFile = "/config/config.yaml"
)

type Config struct {
	Port       string `env:"PORT"`
	ConfigFile string `env:"CONFIG_FILE"`
}

func LoadConfig() (c Config, e error) {
	// Parse environmental variables into Config struct
	err := env.Parse(&c)
	if err != nil {
		return c, err
	}

	// Return the Config struct
	return c, nil
}
