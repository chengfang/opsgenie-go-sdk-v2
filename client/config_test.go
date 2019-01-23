package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateApiKey(t *testing.T) {
	conf := &Config{ApiKey: ""}
	err := conf.Validate()
	assert.Equal(t, err.Error(), "API key cannot be blank.")
}

func TestValidateRetryCount(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.RetryCount = -2
	err := conf.Validate()
	assert.Contains(t, err.Error(), "cannot be less than 1")
}

func TestValidateProxyUrl(t *testing.T) {
	conf := &Config{ApiKey: "an api key"}
	conf.RetryCount = 2
	conf.ProxyUrl = "google.com"
	err := conf.Validate()
	assert.Contains(t, err.Error(), "is not a valid url")
}
