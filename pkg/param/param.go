package param

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/lib"
	"net/url"
	"strconv"
	"strings"
)

type Param struct {
	name   string
	values []string
	err    error
}

func New(name string) *Param {
	return &Param{name, nil, nil}
}

func NewError(name string, err error) (param *Param) {
	param = New(name)
	param.err = err
	return
}

func NewString(name string, value string) (param *Param) {
	param = New(name)
	param.values = []string{value}
	return
}

func NewInt(name string, value int) *Param {
	return NewString(name, strconv.Itoa(value))
}

func NewFloat(name string, value float64) *Param {
	return NewString(name, strconv.FormatFloat(value, 'E', 4, 64))
}

func NewBool(name string, value bool) *Param {
	return NewString(name, strconv.FormatBool(value))
}

func (p *Param) Prepare() error {
	return p.err
}

func (p *Param) Name() string {
	return strings.ToLower(p.name)
}

func (p *Param) Values() ([]string, error) {
	return p.values, p.err
}

func (p *Param) String() string {
	return fmt.Sprintf("%s: %s", p.name, strings.Join(p.values, " "))
}

func (p *Param) Delete(conf *config.Config) (errs []error) {
	if conf == nil {
		conf = config.Default
	}
	if urls, err := p.Values(); lib.AddErr(&errs, err) {
		for _, val := range urls {
			if _, err := url.ParseRequestURI(val); lib.AddErr(&errs, err) {
				lib.AddErr(&errs, lib.RequestDelete(val, conf.HTTPClient))
			}
		}
	}
	return
}
