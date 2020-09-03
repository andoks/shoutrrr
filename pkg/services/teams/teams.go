package teams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service providing teams as a notification service
type Service struct {
	standard.Standard
	config *Config
}

// Send a notification message to Microsoft Teams
func (service *Service) Send(message string, params *types.Params) error {
	body := JSON{
		CardType: "MessageCard",
		Context:  "http://schema.org/extensions",
		Markdown: true,
		Text:     message,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := http.Post(service.config.URL.String(), "application/json", bytes.NewBuffer(jsonBody))
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send notification to teams, response status code %s", res.Status)
	}
	if err != nil {
		return fmt.Errorf(
			"an error occurred while sending notification to teams: %s",
			err.Error(),
		)
	}
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger *log.Logger) error {
	service.Logger.SetLogger(logger)
	service.config = &Config{}
	if err := service.config.SetURL(configURL); err != nil {
		return err
	}

	return nil
}
