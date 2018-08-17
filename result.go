package convertapi

import (
	"encoding/json"
	"github.com/ConvertAPI/convertapi-go/lib"
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
	if resp, err := lib.RespExtractErr(client.PostForm(url, *data)); err == nil {
		defer resp.Body.Close()
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

func (this *Result) Cost() (cost int, err error) {
	<-this.waitCh
	if this.response != nil {
		cost = this.response.ConversionCost
	}
	return cost, this.err
}

func (this *Result) Files() (files []*ResFile, err error) {
	<-this.waitCh
	if this.response != nil {
		files = this.response.Files
	}
	return files, this.err
}

func (this *Result) Urls() (urls []string, err error) {
	files, err := this.Files()
	if err == nil {
		for _, file := range files {
			urls = append(urls, file.Url)
		}
	}
	return
}

func (this *Result) Read(p []byte) (n int, err error) {
	files, err := this.Files()
	if err == nil {
		return files[0].Read(p)
	}
	return
}

func (this *Result) ToFile(file *os.File) (err error) {
	files, err := this.Files()
	if err == nil {
		return files[0].ToFile(file)
	}
	return
}

func (this *Result) ToPath(path string) (files []*os.File, errs []error) {
	if resFiles, err := this.Files(); lib.AddErr(&errs, err) {
		if !lib.IsDir(path) {
			resFiles = []*ResFile{resFiles[0]}
		}

		for _, resFile := range resFiles {
			file, err := resFile.ToPath(path)
			files = append(files, file)
			lib.AddErr(&errs, err)
		}
	}
	return
}

func (this *Result) Delete() (errs []error) {
	if files, err := this.Files(); lib.AddErr(&errs, err) {
		for _, file := range files {
			lib.AddErr(&errs, file.Delete())
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
