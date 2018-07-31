package convertapi_golang

import (
	"fmt"
	"net/url"
	"os"
)

func Convert(fromFormat string, toFormat string, params []*Param, config *Config) (result *Result) {
	result = NewResult()
	go func() {
		if config == nil {
			config = Default
		}
		ignoreParams := []string{"storefile", "async", "jobid", "timeout"}
		values := &url.Values{}
		for _, param := range params {
			if !contains(ignoreParams, param.name) {
				if vals, err := param.Values(); err == nil {
					if len(vals) == 1 {
						values.Add(param.name, vals[0])
					} else {
						for i, val := range vals {
							values.Add(fmt.Sprintf("%s[%d]", param.name, i), val)
						}
					}
				} else {
					result.reject(err)
					return
				}
			}
		}

		query := url.Values{}
		query.Add("secret", config.Secret)
		query.Add("storefile", "true")
		path := fmt.Sprintf("/convert/%s/to/%s?%s", fromFormat, toFormat, query.Encode())
		pathURL, err := url.Parse(path)
		if err != nil {
			result.reject(err)
			return
		}
		convertURL := config.BaseURL.ResolveReference(pathURL)

		result.start(convertURL.String(), values, config.HttpClient)
	}()
	return
}

func ConvertPath(fromPath string, toPath string) (file *os.File, errs []error) {
	res := Convert(pathExt(fromPath), pathExt(toPath), []*Param{
		NewFilePathParam("file", fromPath, nil),
	}, nil)

	if addErr(&errs, res.err) {
		if files, e := res.ToPath(toPath); e == nil {
			file = files[0]
		} else {
			errs = e
		}
	}
	return
}
