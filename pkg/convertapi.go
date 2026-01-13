package convertapi

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"github.com/ConvertAPI/convertapi-go/pkg/lib"
	"github.com/ConvertAPI/convertapi-go/pkg/param"
)

func ConvertDefault(fromFormat string, toFormat string, params ...param.Parameter) (result *Result) {
	return Convert(fromFormat, toFormat, params, nil)
}

func Convert(fromFormat string, toFormat string, params []param.Parameter, conf *config.Config) (result *Result) {
	result = NewResult()
	go func() {
		if conf == nil {
			conf = config.Default
		}
		ignoreParams := []string{"storefile", "async", "jobid"}
		values := &url.Values{}

		paramVals, err := prepareValues(params)
		if err != nil {
			result.reject(err)
			return
		}

		values.Add("storefile", "true")
		for name, vals := range paramVals {
			if !lib.Contains(ignoreParams, name) {
				if len(vals) == 1 {
					values.Add(name, vals[0])
				} else {
					for _, val := range vals {
						values.Add(name, val)
					}
				}
			}
		}

		path := fmt.Sprintf("/v3/convert/%s/to/%s", fromFormat, toFormat)
		pathURL, err := url.Parse(path)
		if err != nil {
			result.reject(err)
			return
		}
		convertURL := conf.BaseURL.ResolveReference(pathURL)

		result.start(convertURL.String(), values, conf.HTTPClient)
	}()
	return
}

func prepareValues(params []param.Parameter) (vals map[string][]string, err error) {
	vals = make(map[string][]string)
	for _, p := range params {
		paramVal, err := p.Values()
		if err != nil {
			return nil, err
		}
		v, ok := vals[p.Name()]
		if ok {
			v = append(v, paramVal...)
		} else {
			v = paramVal
		}
		vals[p.Name()] = v
	}
	return
}

func ConvertPath(fromPath string, toPath string) (file *os.File, errs []error) {
	res := Convert(lib.PathExt(fromPath), lib.PathExt(toPath), []param.Parameter{
		param.NewPath("file", fromPath, nil),
	}, nil)

	if lib.AddErr(&errs, res.err) {
		if files, e := res.ToPath(toPath); e == nil {
			file = files[0]
		} else {
			errs = e
		}
	}
	return
}
