package param

import (
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"sync"
)

type ParamResult struct {
	Param
	res    IResult
	config *config.Config
	sync.Mutex
}

func NewResult(name string, res IResult, conf *config.Config) (param *ParamResult) {
	if conf == nil {
		conf = config.Default
	}
	param = &ParamResult{*New(name), res, conf, sync.Mutex{}}
	return
}

func (this *ParamResult) Prepare() error {
	this.Lock()
	defer this.Unlock()

	if this.values == nil {
		if err := this.Param.Prepare(); err != nil {
			return err
		}

		urls, err := this.res.Urls()
		if err != nil {
			return err
		}
		for _, url := range urls {
			this.values = append(this.values, url)
		}
	}
	return nil
}

func (this *ParamResult) Values() ([]string, error) {
	err := this.Prepare()
	return this.values, err
}
