package convertapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type Result struct {
	waitCh   chan struct{}
	err      error
	response *response
}

type response struct {
	ConversionCost int
	Files          []*ResFile
}

func NewResult() *Result {
	return &Result{make(chan struct{}), nil, nil}
}

func (this *Result) start(url string, data *url.Values, client *http.Client) {
	if resp, err := client.PostForm(url, *data); err == nil {
		response := &response{}
		json.NewDecoder(resp.Body).Decode(response)

		for _, file := range response.Files {
			file.client = client
		}
		this.resolve(response)
	} else {
		this.reject(err)
	}
}

func (this *Result) Cost() int {
	<-this.waitCh
	return this.response.ConversionCost
}

func (this *Result) Files() []*ResFile {
	<-this.waitCh
	return this.response.Files
}

func (this *Result) Read(p []byte) (n int, err error) {
	return this.Files()[0].Read(p)
}

func (this *Result) ToFile(file *os.File) (err error) {
	return this.Files()[0].ToFile(file)
}

func (this *Result) ToPath(path string) (err error) {
	if info, e := os.Stat(path); e == nil && info.IsDir() {
		for _, file := range this.Files() {
			if err = file.ToPath(path); err != nil {
				return
			}
		}
	}
	return this.Files()[0].ToPath(path)
}

func (this *Result) Delete() (errArr []error) {
	for _, file := range this.Files() {
		if err := file.Delete(); err != nil {
			errArr = append(errArr, err)
		}
	}
	return
}

func (this *Result) resolve(response *response) {
	this.response = response
	close(this.waitCh)
}

func (this *Result) reject(err error) {
	this.err = err
	close(this.waitCh)
}
