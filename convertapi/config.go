package convertapi

import (
	"net/http"
	"net/url"
)

var defaultConfig = newDefaultConfig()

type Config struct {
	secret     string
	baseURL    *url.URL
	httpClient *http.Client
}

func newDefaultConfig() *Config {
	url, _ := url.ParseRequestURI("https://v2.convertapi.com")
	return &Config{
		baseURL:    url,
		httpClient: http.DefaultClient,
	}
}

func NewConfig(secret string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = defaultConfig.baseURL
	}
	if transport == nil {
		transport = &http.Transport{}
	}

	return &Config{
		secret:     secret,
		baseURL:    url,
		httpClient: &http.Client{Transport: transport},
	}
}

func Secret(secret string) {
	defaultConfig.secret = secret
}
