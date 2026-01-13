package config

import (
	"fmt"
	"net/http"
	"runtime"
)

type CaTransport struct {
	http.RoundTripper
	AuthCred string
}

func NewCaTransport(authCred string, roundTripper http.RoundTripper) *CaTransport {
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	return &CaTransport{roundTripper, authCred}
}

func (t *CaTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	runtime.Version()
	agent := fmt.Sprintf("ConvertAPI-Go/%d (%s)", Version, runtime.GOOS)
	req.Header.Add("User-Agent", agent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.AuthCred))
	return t.RoundTripper.RoundTrip(req)
}
