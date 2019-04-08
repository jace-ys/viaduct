package config

import (
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/jace-ys/viaduct/pkg/domain"
)

type Config struct {
	Port      string `env:"PORT" envDefault:"80"`
	FilePath  string `env:"FILE_PATH" envDefault:"/config/config.yml"`
	LogPrefix string `env:"LOG_PREFIX" envDefault:"viaduct"`
}

type ServiceDefinitions struct {
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

func RegisterServices(filePath string) (s ServiceDefinitions, e error) {
	// Read the configuration file
	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		return s, err
	}

	// Unmarshal the YAML into the ServiceDefinitions struct
	err = yaml.Unmarshal(out, &s)
	if err != nil {
		return s, err
	}

	// Return the ServiceDefinitions struct
	return s, nil
}
