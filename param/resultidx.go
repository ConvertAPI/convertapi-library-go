package param

import (
	"github.com/ConvertAPI/convertapi-go/config"
)

type ParamResultIdx struct {
	ParamResult
	idx int
}

func NewResultIdx(name string, res IResult, idx int, conf *config.Config) *ParamResultIdx {
	return &ParamResultIdx{*NewResult(name, res, conf), idx}
}

func (this *ParamResultIdx) Values() ([]string, error) {
	err := this.ParamResult.Prepare()
	if this.idx < 0 {
		this.idx = len(this.values) + this.idx
	}
	return []string{this.values[this.idx]}, err
}
