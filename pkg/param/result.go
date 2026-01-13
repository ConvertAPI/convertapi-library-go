package param

import (
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"sync"
)

type ParamResult struct {
	Param
	res    ResultParameter
	config *config.Config
	sync.Mutex
}

func NewResult(name string, res ResultParameter, conf *config.Config) (param *ParamResult) {
	if conf == nil {
		conf = config.Default
	}
	param = &ParamResult{*New(name), res, conf, sync.Mutex{}}
	return
}

func (pr *ParamResult) Prepare() error {
	pr.Lock()
	defer pr.Unlock()

	if pr.values == nil {
		if err := pr.Param.Prepare(); err != nil {
			return err
		}

		ids, err := pr.res.Ids()
		if err != nil {
			return err
		}
		for _, fid := range ids {
			pr.values = append(pr.values, fid)
		}
	}
	return nil
}

func (pr *ParamResult) Values() ([]string, error) {
	err := pr.Prepare()
	return pr.values, err
}
