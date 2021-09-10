package config

import (
	"net/http"
	"net/url"
)

var Default = newDefaultConfig()

type Config struct {
	Secret     string
	BaseURL    *url.URL
	HttpClient *http.Client
	Token      string
	ApiKey     string
}

func newDefaultConfig() *Config {
	baseUrl, _ := url.ParseRequestURI("https://v2.convertapi.com")
	client := &http.Client{Transport: NewCaTransport(nil)}
	return &Config{"", baseUrl, client, "", ""}
}

func NewConfig(secret string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = Default.BaseURL
	}

	caTransport := NewCaTransport(transport)
	return &Config{secret, url, &http.Client{Transport: caTransport}, "", ""}
}

func NewConfigToken(token string, apikey string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = Default.BaseURL
	}

	caTransport := NewCaTransport(transport)
	return &Config{"", url, &http.Client{Transport: caTransport}, token, apikey}
}

func (c Config) AddAuth(query url.Values) {
	if c.Secret == "" {
		query.Add("token", c.Token)
		query.Add("apikey", c.ApiKey)
	} else {
		query.Add("secret", c.Secret)
	}
}
