package convertapi

import (
	"fmt"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/lib"
	"github.com/ConvertAPI/convertapi-go/param"
	"net/url"
	"os"
)

func ConvDef(fromFormat string, toFormat string, params ...param.IParam) (result *Result) {
	return Convert(fromFormat, toFormat, params, nil)
}

func Convert(fromFormat string, toFormat string, params []param.IParam, conf *config.Config) (result *Result) {
	result = NewResult()
	go func() {
		if conf == nil {
			conf = config.Default
		}
		ignoreParams := []string{"storefile", "async", "jobid", "timeout"}
		values := &url.Values{}

		paramVals, err := prepareValues(params)
		if err != nil {
			result.reject(err)
			return
		}

		for name, vals := range paramVals {
			if !lib.Contains(ignoreParams, name) {
				if len(vals) == 1 {
					values.Add(name, vals[0])
				} else {
					for i, val := range vals {
						values.Add(fmt.Sprintf("%s[%d]", name, i), val)
					}
				}
			}
		}

		query := url.Values{}
		query.Add("secret", conf.Secret)
		query.Add("storefile", "true")
		path := fmt.Sprintf("/convert/%s/to/%s?%s", fromFormat, toFormat, query.Encode())
		pathURL, err := url.Parse(path)
		if err != nil {
			result.reject(err)
			return
		}
		convertURL := conf.BaseURL.ResolveReference(pathURL)

		result.start(convertURL.String(), values, conf.HttpClient)
	}()
	return
}

func prepareValues(params []param.IParam) (vals map[string][]string, err error) {
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
	res := Convert(lib.PathExt(fromPath), lib.PathExt(toPath), []param.IParam{
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
