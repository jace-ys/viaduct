package config

import (
	"io/ioutil"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"

	"github.com/jace-ys/viaduct/pkg/domain"
)

var (
	DefaultPort       = "80"
	DefaultConfigFile = "/config/config.yml"
)

type Config struct {
	Port       string `env:"PORT"`
	ConfigFile string `env:"CONFIG_FILE"`
}

type ServiceRegistry struct {
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

func RegisterServices(configFile string) (s ServiceRegistry, e error) {
	// Read the configuration file
	out, err := ioutil.ReadFile(configFile)
	if err != nil {
		return s, err
	}

	// Unmarshal the YAML into the ServiceRegistry struct
	err = yaml.Unmarshal(out, &s)
	if err != nil {
		return s, err
	}

	// Return the ServiceRegistry struct
	return s, nil
}
