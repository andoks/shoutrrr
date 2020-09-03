package teams

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
)

// Config for use within the teams plugin
type Config struct {
	standard.QuerylessConfig
	standard.EnumlessConfig
	URL url.URL
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	return &config.URL
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error {
	url.Scheme = "https"
	config.URL = *url
	return nil
}

// CreateConfigFromURL for use within the teams plugin
func (service *Service) CreateConfigFromURL(url *url.URL) (*Config, error) {
	config := Config{}
	err := config.SetURL(url)
	return &config, err
}

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "teams+https"
)
