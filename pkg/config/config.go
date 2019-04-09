package config

import (
	"io/ioutil"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"

	"github.com/jace-ys/viaduct/pkg/domain"
)

type Config struct {
	Port       string `env:"PORT" envDefault:"80"`
	ConfigPath string `env:"CONFIG_PATH" envDefault:"/config/config.yml"`
	LogPrefix  string `env:"LOG_PREFIX" envDefault:"viaduct"`
}

type ServiceRegister struct {
	Services map[string]domain.Service
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

func RegisterServices(configPath string) (s ServiceRegister, e error) {
	// Read the configuration file
	out, err := ioutil.ReadFile(configPath)
	if err != nil {
		return s, err
	}

	// Unmarshal the YAML into the ServiceRegister struct
	err = yaml.Unmarshal(out, &s)
	if err != nil {
		return s, err
	}

	// Return the ServiceRegister struct
	return s, nil
}
