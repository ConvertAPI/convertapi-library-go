package convertapi_golang

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Param struct {
	client *http.Client
	waitCh chan struct{}
	err    error
	name   string
	value  []string
}

func newParam(name string) *Param {
	return &Param{nil, make(chan struct{}), nil, strings.ToLower(name), nil}
}

func newParamWithValueResolved(name string, value string) (param *Param) {
	param = newParam(name)
	param.resolve([]string{value})
	return
}

func NewStringParam(name string, value string) *Param {
	return newParamWithValueResolved(name, value)
}

func NewIntParam(name string, value int) *Param {
	return newParamWithValueResolved(name, strconv.Itoa(value))
}

func NewFloatParam(name string, value float64) *Param {
	return newParamWithValueResolved(name, strconv.FormatFloat(value, 'E', 4, 64))
}

func NewBoolParam(name string, value bool) *Param {
	return newParamWithValueResolved(name, strconv.FormatBool(value))
}

func NewReaderParam(name string, value io.Reader, filename string, config *Config) (param *Param) {
	if config == nil {
		config = Default
	}
	param = newParam(name)
	param.client = config.HttpClient

	go func() {
		query := url.Values{}
		query.Add("filename", filename)

		pathURL, err := url.Parse("/upload?" + query.Encode())
		if err != nil {
			param.reject(err)
			return
		}

		uploadURL := config.BaseURL.ResolveReference(pathURL)
		resp, err := config.HttpClient.Post(uploadURL.String(), "application/octet-stream", value)
		defer resp.Body.Close()
		if err != nil {
			param.reject(err)
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		param.resolve([]string{buf.String()})
	}()

	return
}

func NewFileParam(name string, value *os.File, config *Config) *Param {
	return NewReaderParam(name, value, filepath.Base(value.Name()), config)
}

func NewFilePathParam(name string, value string, config *Config) *Param {
	if f, err := os.Open(value); err == nil {
		return NewFileParam(name, f, config)
	} else {
		param := newParam(name)
		param.reject(err)
		return param
	}
}

func NewResultParam(name string, value *Result, config *Config) (param *Param) {
	if config == nil {
		config = Default
	}
	param = newParam(name)
	param.client = config.HttpClient

	go func() {
		files, err := value.Files()
		if err == nil {
			var paths []string
			for _, file := range files {
				paths = append(paths, file.Url)
			}
			param.resolve(paths)
		} else {
			param.reject(err)
		}
	}()

	return
}

func (this *Param) Values() (value []string, err error) {
	<-this.waitCh
	return this.value, this.err
}

func (this *Param) Delete() (errs []error) {
	if this.client != nil {
		if vals, err := this.Values(); addErr(&errs, err) {
			for _, val := range vals {
				if _, err = url.ParseRequestURI(val); addErr(&errs, err) {
					addErr(&errs, requestDelete(val, this.client))
				}
			}
		}
	}
	return
}

func (this *Param) resolve(value []string) {
	this.value = value
	close(this.waitCh)
}

func (this *Param) reject(err error) {
	this.err = err
	close(this.waitCh)
}
