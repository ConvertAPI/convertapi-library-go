package param

import (
	"github.com/ConvertAPI/convertapi-go/config"
)

type IParam interface {
	Prepare() error
	Name() string
	Values() ([]string, error)
	Delete(conf *config.Config) []error
}

type IResult interface {
	Urls() ([]string, error)
}
