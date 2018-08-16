package config

import (
	"fmt"
	"net/http"
	"runtime"
)

type CaTransport struct {
	http.RoundTripper
}

func NewCaTransport(roundTripper http.RoundTripper) *CaTransport {
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	return &CaTransport{roundTripper}
}

func (this *CaTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	runtime.Version()
	agent := fmt.Sprintf("convertapi-go/%d (%s)", Version, runtime.GOOS)
	req.Header.Add("User-Agent", agent)
	return this.RoundTripper.RoundTrip(req)
}
