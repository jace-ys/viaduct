package domain

import (
	"net/url"
	"strings"
)

type Service struct {
	Name             string
	Prefix           string
	UpstreamUrl      *ServiceURL `yaml:"upstream_url"`
	Methods          []string
	AllowCrossOrigin bool `yaml:"allow_cross_origin"`
}

type ServiceURL struct {
	*url.URL
}

// Tell yaml parser how to unmarshal a string to type Url
func (u *ServiceURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
