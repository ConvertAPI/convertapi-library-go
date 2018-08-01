package convertapi

import (
	"net/http"
	"net/url"
)

var Default = newDefaultConfig()

type Config struct {
	Secret     string
	BaseURL    *url.URL
	HttpClient *http.Client
}

func newDefaultConfig() *Config {
	baseUrl, _ := url.ParseRequestURI("https://v2.convertapi.com")
	client := &http.Client{Transport: NewCaTransport(nil)}
	return &Config{"", baseUrl, client}
}

func NewConfig(secret string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = Default.BaseURL
	}

	caTransport := NewCaTransport(transport)
	return &Config{secret, url, &http.Client{Transport: caTransport}}
}
