package config

import (
	"net/http"
	"net/url"
)

var Default *Config

type Config struct {
	BaseURL     *url.URL
	CaTransport *CaTransport
	HttpClient  *http.Client
}

func NewDefault(authCred string) *Config {
	baseUrl, _ := url.ParseRequestURI("https://v2.convertapi.com")
	transport := NewCaTransport(authCred, nil)
	client := &http.Client{Transport: transport}
	return &Config{baseUrl, transport, client}
}

func New(authCred string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = Default.BaseURL
	}

	caTransport := NewCaTransport(authCred, transport)
	return &Config{BaseURL: url, HttpClient: &http.Client{Transport: caTransport}}
}
