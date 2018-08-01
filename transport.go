package convertapi

import "net/http"

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
	req.Header.Add("User-Agent", "convertapi-go")
	return this.RoundTripper.RoundTrip(req)
}
