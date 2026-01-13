package param

import (
	"bytes"
	"fmt"
	"github.com/ConvertAPI/convertapi-go/pkg/config"
	"io"
	"net/url"
	"strings"
	"sync"
)

type ParamReader struct {
	Param
	reader   io.Reader
	fileName string
	config   *config.Config
	sync.Mutex
}

func NewReader(name string, reader io.Reader, filename string, conf *config.Config) *ParamReader {
	if conf == nil {
		conf = config.Default
	}
	return &ParamReader{*New(name), reader, filename, conf, sync.Mutex{}}
}

func (pr *ParamReader) Prepare() error {
	pr.Lock()
	defer pr.Unlock()

	if pr.values == nil {
		if err := pr.Param.Prepare(); err != nil {
			return err
		}

		query := url.Values{}
		query.Add("filename", pr.fileName)

		pathURL, err := url.Parse("/v3/upload?" + query.Encode())
		if err != nil {
			return err
		}

		uploadURL := pr.config.BaseURL.ResolveReference(pathURL)
		resp, err := pr.config.HTTPClient.Post(uploadURL.String(), "application/octet-stream", pr.reader)

		if err != nil {
			return err
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return err
		}
		pr.values = []string{buf.String()}
	}
	return nil
}

func (pr *ParamReader) Values() ([]string, error) {
	err := pr.Prepare()
	return pr.values, err
}

func (pr *ParamReader) String() string {
	return fmt.Sprintf("%s: %s -> %s", pr.name, pr.fileName, strings.Join(pr.values, " "))
}
