package api

import (
	"io/ioutil"
	"net/url"
	"strings"

	"gopkg.in/yaml.v2"
)

type Registry struct {
	APIs map[string]API
}

func RegisterAPIs(configFile string) (ar Registry, e error) {
	// Read the configuration file
	out, err := ioutil.ReadFile(configFile)
	if err != nil {
		return ar, err
	}

	// Unmarshal the .yaml file into the ApiRegistry struct
	err = yaml.Unmarshal(out, &ar)
	if err != nil {
		return ar, err
	}

	// Return the ApiRegistry struct
	return ar, nil
}

type API struct {
	Name             string
	Prefix           string
	UpstreamUrl      *Url `yaml:"upstream_url"`
	Methods          []string
	AllowCrossOrigin bool `yaml:"allow_cross_origin"`
	Middlewares      map[string]bool
}

type Url struct {
	*url.URL
}

// Tell yaml parser how to unmarshal a string to type Url
func (u *Url) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	// Unmarshal the value into a string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	// Parse the URL string into a url.URL type
	target, err := url.Parse(strings.TrimSuffix(s, "/"))
	if err != nil {
		return err
	}

	// Set the anonymous field on u
	u.URL = target

	// Return nil error
	return nil
}
