package convertapi

import (
	"encoding/json"
	"github.com/ConvertAPI/convertapi-go/pkg/lib"
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

func (r *Result) start(url string, data *url.Values, client *http.Client) {
	if resp, err := lib.RespExtractErr(client.PostForm(url, *data)); err == nil {
		defer resp.Body.Close()
		response := &response{}
		json.NewDecoder(resp.Body).Decode(response)

		for _, file := range response.Files {
			file.client = client
		}
		r.resolve(response)
	} else {
		r.reject(err)
	}
}

func (r *Result) Cost() (cost int, err error) {
	<-r.waitCh
	if r.response != nil {
		cost = r.response.ConversionCost
	}
	return cost, r.err
}

func (r *Result) Files() (files []*ResFile, err error) {
	<-r.waitCh
	if r.response != nil {
		files = r.response.Files
	}
	return files, r.err
}

func (r *Result) Ids() (ids []string, err error) {
	files, err := r.Files()
	if err == nil {
		for _, file := range files {
			ids = append(ids, file.FileID)
		}
	}
	return
}

func (r *Result) Urls() (urls []string, err error) {
	files, err := r.Files()
	if err == nil {
		for _, file := range files {
			urls = append(urls, file.URL)
		}
	}
	return
}

func (r *Result) Read(p []byte) (n int, err error) {
	files, err := r.Files()
	if err == nil {
		return files[0].Read(p)
	}
	return
}

func (r *Result) ToFile(file *os.File) (err error) {
	files, err := r.Files()
	if err == nil {
		return files[0].ToFile(file)
	}
	return
}

func (r *Result) ToPath(path string) (files []*os.File, errs []error) {
	if resFiles, err := r.Files(); lib.AddErr(&errs, err) {
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

func (r *Result) Delete() (errs []error) {
	if files, err := r.Files(); lib.AddErr(&errs, err) {
		for _, file := range files {
			lib.AddErr(&errs, file.Delete())
		}
	}
	return
}

func (r *Result) resolve(response *response) {
	r.response = response
	close(r.waitCh)
}

func (r *Result) reject(err error) {
	r.err = err
	close(r.waitCh)
}
