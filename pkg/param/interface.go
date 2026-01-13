package param

import (
	"github.com/ConvertAPI/convertapi-go/pkg/config"
)

type Parameter interface {
	Prepare() error
	Name() string
	Values() ([]string, error)
	Delete(conf *config.Config) []error
}

type ResultParameter interface {
	Ids() ([]string, error)
}
