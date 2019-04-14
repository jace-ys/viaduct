package config

import (
	"io/ioutil"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"

	"github.com/jace-ys/viaduct/pkg/domain"
)

var (
	DefaultPort       = "80"
	DefaultConfigFile = "/config/config.yaml"
)

type Config struct {
	Port       string `env:"PORT"`
	ConfigFile string `env:"CONFIG_FILE"`
}

type ApiRegistry struct {
	Apis map[string]domain.Api
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

func RegisterApis(configFile string) (ar ApiRegistry, e error) {
	// Read the configuration file
	out, err := ioutil.ReadFile(configFile)
	if err != nil {
		return ar, err
	}

	// Unmarshal the YAML into the ApiRegistry struct
	err = yaml.Unmarshal(out, &ar)
	if err != nil {
		return ar, err
	}

	// Return the ApiRegistry struct
	return ar, nil
}
