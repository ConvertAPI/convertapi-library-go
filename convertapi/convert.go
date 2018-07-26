package convertapi

import (
	"fmt"
	"net/url"
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
				if val, err := param.Value(); err == nil {
					values.Add(param.name, val)
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
