package param

import (
	"github.com/ConvertAPI/convertapi-go/pkg/config"
)

type ParamResultIdx struct {
	ParamResult
	idx int
}

func NewResultIdx(name string, res ResultParameter, idx int, conf *config.Config) *ParamResultIdx {
	return &ParamResultIdx{*NewResult(name, res, conf), idx}
}

func (pri *ParamResultIdx) Values() ([]string, error) {
	err := pri.ParamResult.Prepare()
	if pri.idx < 0 {
		pri.idx = len(pri.values) + pri.idx
	}
	return []string{pri.values[pri.idx]}, err
}
