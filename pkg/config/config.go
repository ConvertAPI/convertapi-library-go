package config

import (
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://v2.convertapi.com"

var Default *Config

type Config struct {
	BaseURL     *url.URL
	CaTransport *CaTransport
	HttpClient  *http.Client
}

func NewDefault(authCred string) *Config {
	baseUrl, _ := url.ParseRequestURI(defaultBaseURL)
	transport := NewCaTransport(authCred, nil)
	client := &http.Client{Transport: transport}
	return &Config{baseUrl, transport, client}
}

// New returns a configuration for another API domain (url) or HTTP transport (e.g. a proxy). A nil url
// keeps the default configuration's domain, and a nil transport uses http.DefaultTransport.
func New(authCred string, url *url.URL, transport *http.Transport) *Config {
	if url == nil {
		if Default != nil {
			url = Default.BaseURL
		} else {
			url = NewDefault(authCred).BaseURL
		}
	}

	// A nil *http.Transport inside the RoundTripper interface is not a nil interface, so pass nil itself.
	var roundTripper http.RoundTripper
	if transport != nil {
		roundTripper = transport
	}
	caTransport := NewCaTransport(authCred, roundTripper)
	return &Config{BaseURL: url, CaTransport: caTransport, HttpClient: &http.Client{Transport: caTransport}}
}
