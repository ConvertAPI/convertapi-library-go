package config

import (
	"net/http"
	"net/url"
)

var Default *Config

type Config struct {
	BaseURL     *url.URL
	CaTransport *CaTransport
	HTTPClient  *http.Client
}

func NewDefault(authCred string) *Config {
	baseURL, _ := url.ParseRequestURI("https://api.convertapi.io")
	transport := NewCaTransport(authCred, nil)
	client := &http.Client{Transport: transport}
	return &Config{baseURL, transport, client}
}

func New(authCred string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		url = Default.BaseURL
	}

	caTransport := NewCaTransport(authCred, transport)
	return &Config{BaseURL: url, HTTPClient: &http.Client{Transport: caTransport}}
}
