package convertapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Result struct {
	waitCh   chan struct{}
	err      error
	response *Response
}

type Response struct {
	ConversionCost int
	Files          []File
}

type File struct {
	FileName string
	FileSize int
	Url      string
}

func NewResult() *Result {
	return &Result{make(chan struct{}), nil, nil}
}

func (this *Result) start(url string, data *url.Values, client *http.Client) {
	if resp, err := client.PostForm(url, *data); err == nil {
		response := &Response{}
		json.NewDecoder(resp.Body).Decode(response)
		this.resolve(response)
	} else {
		this.reject(err)
	}
}

func (this *Result) Response() (*Response, error) {
	<-this.waitCh
	return this.response, this.err
}

func (this *Result) resolve(response *Response) {
	this.response = response
	close(this.waitCh)
}

func (this *Result) reject(err error) {
	this.err = err
	close(this.waitCh)
}
