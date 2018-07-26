package convertapi

import (
	"bytes"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Param struct {
	waitCh chan struct{}
	err    error
	name   string
	value  []string
}

func newParam(name string) *Param {
	return &Param{make(chan struct{}), nil, strings.ToLower(name), nil}
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
		config = defaultConfig
	}
	param = newParam(name)

	go func() {
		query := url.Values{}
		query.Add("filename", filename)

		pathURL, err := url.Parse("/upload?" + query.Encode())
		if err != nil {
			param.reject(err)
			return
		}

		uploadURL := config.baseURL.ResolveReference(pathURL)
		resp, err := config.httpClient.Post(uploadURL.String(), "application/octet-stream", value)
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
	return NewReaderParam(name, value, value.Name(), config)
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

func (this *Param) Value() (value string, err error) {
	<-this.waitCh
	return this.value[0], this.err
}

func (this *Param) Delete(config *Config) (finishCh chan struct{}) {
	finishCh = make(chan struct{})
	go func() {
		defer close(finishCh)
		if config == nil {
			config = defaultConfig
		}
		if val, err := this.Value(); err == nil {
			if _, err := url.ParseRequestURI(val); err == nil {
				requestDelete(val, config.httpClient)
			}
		}
	}()
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
